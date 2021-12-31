package javasrc;

import java.lang.reflect.Type;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonSerializationContext;
import com.google.gson.JsonSerializer;

import spoon.reflect.declaration.CtPackage;

public class CtPackageJSONSerializer implements JsonSerializer<CtPackage>{
  @Override
    public JsonElement serialize(CtPackage pkg, Type arg1, JsonSerializationContext ctx) {
        var o = new JsonObject();
        o.addProperty("qualified_name",pkg.getQualifiedName());
        o.addProperty("class", pkg.getClass().getName());
        o.addProperty("simpleName", pkg.getSimpleName());
        o.add("reference", ctx.serialize(pkg.getReference()));
        o.add("children_packages", ctx.serialize(pkg.getPackages()));
        o.add("count_children_packages", ctx.serialize(pkg.getPackages().size()));
        o.add("types", ctx.serialize(pkg.getTypes()));
        return o;
    }
   
}
