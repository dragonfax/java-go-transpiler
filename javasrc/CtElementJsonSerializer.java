package javasrc;

import java.lang.reflect.Type;

import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonSerializationContext;
import com.google.gson.JsonSerializer;

import spoon.reflect.declaration.CtElement;

public class CtElementJsonSerializer implements JsonSerializer<CtElement>{

    @Override
    public JsonElement serialize(CtElement element, Type arg1, JsonSerializationContext ctx) {
        var o = new JsonObject();
        o.add("metadata",ctx.serialize(element.getAllMetadata()));
        return o;
    }

}
