package javasrc;

import java.lang.reflect.Type;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonSerializationContext;
import com.google.gson.JsonSerializer;

import spoon.reflect.declaration.CtPackage;
import spoon.reflect.declaration.CtType;


public class CtTypeJsonSerializer implements JsonSerializer<CtType<?>>{
  @Override
    public JsonElement serialize(CtType<?> pkg, Type arg1, JsonSerializationContext ctx) {
        var o = new JsonObject();
        o.addProperty("qualified_name",pkg.getQualifiedName());
        o.addProperty("class", pkg.getClass().getName());
        o.addProperty("simpleName", pkg.getSimpleName());
        o.add("reference", ctx.serialize(pkg.getReference()));
        o.add("fields", ctx.serialize(pkg.getFields()));
        return o;
    }
 
}
