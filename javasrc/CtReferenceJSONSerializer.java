package javasrc;

import spoon.reflect.declaration.CtModule;
import spoon.reflect.reference.CtReference;
import com.google.gson.JsonObject;
import com.google.gson.JsonSerializationContext;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonElement;

import java.lang.reflect.Type;

import com.google.gson.Gson;
import com.google.gson.JsonSerializer;


public class CtReferenceJSONSerializer implements JsonSerializer<CtReference>{

    @Override
    public JsonElement serialize(CtReference ref, Type arg1, JsonSerializationContext arg2) {
        var o = new JsonObject();
        o.addProperty("class", ref.getClass().getName());
        o.addProperty("refValue", ref.getSimpleName());
        return o;
    }
    
}
