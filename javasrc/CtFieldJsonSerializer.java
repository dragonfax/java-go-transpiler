package javasrc;

import java.lang.reflect.Type;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonSerializationContext;
import com.google.gson.JsonSerializer;

import spoon.reflect.declaration.CtField;
import spoon.reflect.declaration.CtPackage;
import spoon.reflect.declaration.CtType;


public class CtFieldJsonSerializer implements JsonSerializer<CtField<?>>{

    @Override
    public JsonElement serialize(CtField<?> field, Type arg1, JsonSerializationContext ctx) {
        var o = new JsonObject();
        o.addProperty("class", field.getClass().getName());
        o.addProperty("simpleName", field.getSimpleName());
        o.add("reference", ctx.serialize(field.getReference()));
        // o.add("assignemnt", ctx.serialize(field.getAssignment()));
        return o;
    }

}
