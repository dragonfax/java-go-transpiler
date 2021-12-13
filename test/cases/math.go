package entities

import "github.com/dragonfax/delver_converted/com/badlogic/gdx/math"
import "github.com/dragonfax/delver_converted/com/interrupt/dungeoneer"
import "github.com/dragonfax/delver_converted/com/interrupt/dungeoneer/annotations"
import "github.com/dragonfax/delver_converted/com/interrupt/dungeoneer/game"
import "github.com/dragonfax/delver_converted/com/interrupt/dungeoneer/game/Level"
import "github.com/dragonfax/delver_converted/com/interrupt/dungeoneer/gfx"

type DynamicLight struct {
	*Entity

	LightColor Vector3

	Range float64

	LightType LightType

	hasHalo bool

	haloMode HaloMode

	on bool

	toggleAnimationTime float64
	toggleLerpTime float64

	haloOffset float64
	haloSize float64

	haloSizeMod float64

	time float64
	workColor Vector3
	workRange float64
	colorLerpTarget Vector3
	rangeLerpTarget float64
	lerpTimer float64
	lerpTime float64
	killAfterLerp *bool

	workVector3 Vector3

}

func NewDynamicLight() *DynamicLight {

	this := &DynamicLight{
		Entity: NewEntity(),

		LightColor: NewVector3(1,1,1),

		Range: 3.2,

		LightType: steady,

		hasHalo : false,

		haloMode : NONE,

		on : true,

		toggleAnimationTime : 0.0,
		toggleLerpTime : 1.0,

		haloOffset : 0.25,
		haloSize : 0.5,

		haloSizeMod : 1.0,

		time : 0.0,
		workColor : NewVector3(),
		workRange : 3.2,
		colorLerpTarget : nil,
		rangeLerpTarget : 3.2,
		lerpTimer : 0.0,
		lerpTime : 0.0,
		killAfterLerp : nil,

		workVector3 : NewVector3(),
	}

	this.hidden = true
	this.spriteAtlas = "editor"
	this.tex = 12
	this.isSolid = false

	return this
}

	
type LightType int

const (
	steady LightType = iota 
	fire
	flicker_on
	flicker_off
	sin_slow
	sin_fast
	torch
	sin_slight
)


func NewDynamicLight4(x, y, z float64, lightColor Vector3) {
	this := NewDynamicLight()

	// TODO  shouldn't be doing NewEntity twice, and throwing away the old one.
	this.Entity = NewEntity4(x, y, 0, false)

	this.z = z;
	this.artType = hidden;
	this.lightColor = lightColor;
}

func NewDynamicLight5(float x, float y, float z, float range, Vector3 lightColor) {
	super(x, y, 0, false);
	this.z = z;
	this.range = range;
	artType = ArtType.hidden;
	this.lightColor = lightColor;
}

func (this *DynamicLight) updateLightColor(delta float64) {
	this.time += delta;
	
	this.workColor.set(lightColor)
	this.workRange = this.Range

	if this.toggleAnimationTime != 0 {
		// animate when turning on / off
		if this.on && this.toggleLerpTime < 1 {
			this.toggleLerpTime += delta / this.toggleAnimationTime
		} else if !this.on && this.toggleLerpTime > 0 {
			this.toggleLerpTime -= delta / this.toggleAnimationTime;
		}

		// clamp!
		if this.toggleLerpTime < 0 {
			this.toggleLerpTime = 0
		}
		if this.toggleLerpTime > 1 {
			this.toggleLerpTime = 1
		}
	}
	
	if this.lightType == steady {
		// steady lights do nothing
	} else if this.lightType == fire {
		this.workColor.scl(1 - math.Sin(this.time * 0.11) * 0.1)
		this.workColor.scl(1 - math.Sin(this.time * 0.147) * 0.1)
		this.workColor.scl(1 - math.Sin(this.time * 0.263) * 0.1)
		
		this.workRange *= 1 - math.Sin(this.time * 0.111) * 0.05
		this.workRange *= 1 - math.Sin(this.time * 0.1477) * 0.05
		this.workRange *= 1 - math.Sin(this.time * 0.2631) * 0.05
	} else if this.lightType == torch {
		this.workColor.scl(1 - math.Sin(this.time * 0.11) * 0.5)
		this.workColor.scl(1 - math.Sin(this.time * 0.147) * 0.5)
		this.workColor.scl(1 - math.Sin(this.time * 0.263) * 0.5)
		
		this.workRange *= 1 - math.Sin(this.time * 0.111) * 0.05
		this.workRange *= 1 - math.Sin(this.time * 0.1477) * 0.05
		this.workRange *= 1 - math.Sin(this.time * 0.2631) * 0.05
	} else if this.lightType == flicker_on {
		a := Game.Rand.nextFloat()
		if a > 0.95 {
			b := 1.0
		} else {
			b := 0.0
		}
		this.workColor.scl(b)
	} else if this.lightType == flicker_off {
		this.workColor.scl(ternaryFloat64(Game.rand.nextFloat() > 0.95, 0.0, 1.0));
	} else if this.lightType == sin_slow {
		this.workColor.scl(math.Sin(this.time * 0.02) + 1)
	} else if this.lightType == sin_slight {
		this.workColor.scl((math.Sin(this.time * 0.05) + 1) * 0.2 + 1)
	} else if this.lightType == sin_fast {
		this.workColor.scl(math.Sin(this.time * 0.2) + 1)
	}

	if this.toggleLerpTime > 0 && this.toggleLerpTime < 1 {
		this.workColor.scl(Interpolation.linear.apply(this.toggleLerpTime))
	}
	
	if this.colorLerpTarget != nil {
		lerpA := this.lerpTimer / this.lerpTime
		this.workColor.lerp(this.colorLerpTarget, lerpA)
		this.workRange = Interpolation.linear.apply(this.Range, this.rangeLerpTarget, lerpA)
		this.lerpTimer += delta;
		
		if this.lerpTimer >= this.lerpTime {
			this.workColor.set(this.colorLerpTarget)
			
			this.colorLerpTarget = nil
			
			if this.killAfterLerp != nil && this.killAfterLerp {
				this.isActive = false
			}
		}
	}

	if math.IsNaN(this.workRange)) {
		this.workRange = 0
	}

	this.haloSize = this.workRange * 0.175;
	this.haloSize *= Interpolation.circleOut.apply(len(this.workColor) * 0.4);
	this.haloSize *= this.haloSizeMod;
}

func (this *DynamicLight) Tick(level Level, delta float64) {
	if !GameManager.Renderer.EnableLighting {
		return
	}

	this.UpdateLightColor(delta)
	
	if this.IsActive && (this.On || (this.ToggleLerpTime > 0 && this.ToggleLerpTime < 1)) {
		if Game.Instance.Camera == nil || Game.Instance.Camera.Frustum.SphereInFrustum(this.WorkVector3.Set(this.X,this.Z,this.Y), this.Range * 1.5) {
			light := GlRenderer.GetLight()
			if light != nil {
				light.Color.Set(this.WorkColor.X, this.WorkColor.Y, this.WorkColor.Z)
				light.Position.Set(this.X, this.Z, this.Y)
				light.Range = this.WorkRange
			}
		}
	}
}

func (this *DynamicLight) EditorTick(level Level, delta float64) {
	this.Entity.EditorTick(level, delta)
	this.Tick(level, delta)
}

func (this *DynamicLight) StartLerp(endColor Vector3, time float64, killAfter bool) *DynamicLight {
	this.ColorLerpTarget = endColor
	this.RangeLerpTarget = range
	this.KillAfterLerp = killAfter
	
	this.LerpTime = time
	this.LerpTimer = 0.0
	
	return this
}

func (this *DynamicLight) StartLerp(endColor Vector3, endRange float64, time float64, killAfter bool) *DynamicLight {
	this.ColorLerpTarget = endColor
	this.RangeLerpTarget = endRange
	this.KillAfterLerp = killAfter
	
	this.LerpTime = time
	this.LerpTimer = 0.0
	
	return this
}

func (this *DynamicLight) SetHaloMode(haloMode HaloMode) *DynamicLight {
	this.haloMode = haloMode
	return this
}

func (this *DynamicLight) Init(level *Level, source *Source) {
	this.Entity.Init(level, source)
	this.time = game.Rand.NextFloat() * 5000.0

	if this.on )  {
		this.ToggleLerpTime = 1
	} else {
		this.ToggleLerpTime = 0
	}

}

func (this *DynamicLight) OnTrigger(instigator *Entity, value string) {
	this.on = !this.on
}

func (this *DynamicLight) GetHaloMode() HaloMode {
	return this.haloMode
}