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

void updateLightColor(float delta) {
	time += delta;
	
	workColor.set(lightColor);
	workRange = range;

	if(toggleAnimationTime != 0) {
		// animate when turning on / off
		if(on && toggleLerpTime < 1)
			toggleLerpTime += delta / toggleAnimationTime;
		else if(!on && toggleLerpTime > 0)
			toggleLerpTime -= delta / toggleAnimationTime;

		// clamp!
		if(toggleLerpTime < 0)
			toggleLerpTime = 0;
		if(toggleLerpTime > 1)
			toggleLerpTime = 1;
	}
	
	if(lightType == LightType.steady) {
		// steady lights do nothing
	}
	else if(lightType == LightType.fire) {
		workColor.scl(1 - (float)Math.sin(time * 0.11f) * 0.1f);
		workColor.scl(1 - (float)Math.sin(time * 0.147f) * 0.1f);
		workColor.scl(1 - (float)Math.sin(time * 0.263f) * 0.1f);
		
		workRange *= 1 - (float)Math.sin(time * 0.111f) * 0.05f;
		workRange *= 1 - (float)Math.sin(time * 0.1477f) * 0.05f;
		workRange *= 1 - (float)Math.sin(time * 0.2631f) * 0.05f;
	}
	else if(lightType == LightType.torch) {
		workColor.scl(1 - (float)Math.sin(time * 0.11f) * 0.5f);
		workColor.scl(1 - (float)Math.sin(time * 0.147f) * 0.5f);
		workColor.scl(1 - (float)Math.sin(time * 0.263f) * 0.5f);
		
		workRange *= 1 - (float)Math.sin(time * 0.111f) * 0.05f;
		workRange *= 1 - (float)Math.sin(time * 0.1477f) * 0.05f;
		workRange *= 1 - (float)Math.sin(time * 0.2631f) * 0.05f;
	}
	else if(lightType == LightType.flicker_on) {
		workColor.scl(Game.rand.nextFloat() > 0.95f ? 1f : 0f);
	}
	else if(lightType == LightType.flicker_off) {
		workColor.scl(Game.rand.nextFloat() > 0.95f ? 0f : 1f);
	}
	else if(lightType == LightType.sin_slow) {
		workColor.scl((float)Math.sin(time * 0.02f) + 1f);
	}
	else if(lightType == LightType.sin_slight) {
		workColor.scl((float)(Math.sin(time * 0.05f) + 1f) * 0.2f + 1f);
	}
	else if(lightType == LightType.sin_fast) {
		workColor.scl((float)Math.sin(time * 0.2f) + 1f);
	}

	if(toggleLerpTime > 0 && toggleLerpTime < 1) {
		workColor.scl(Interpolation.linear.apply(toggleLerpTime));
	}
	
	if(colorLerpTarget != null) {
		float lerpA = lerpTimer / lerpTime;
		workColor.lerp(colorLerpTarget, lerpA);
		workRange = Interpolation.linear.apply(range, rangeLerpTarget, lerpA);
		lerpTimer += delta;
		
		if(lerpTimer >= lerpTime) {
			workColor.set(colorLerpTarget);
			
			colorLerpTarget = null;
			
			if(killAfterLerp != null && killAfterLerp) isActive = false;
		}
	}

	if(Float.isNaN(workRange)) workRange = 0;

	haloSize = workRange * 0.175f;
	haloSize *= Interpolation.circleOut.apply(workColor.len() * 0.4f);
	haloSize *= haloSizeMod;
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