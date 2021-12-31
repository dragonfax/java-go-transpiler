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


public class NullJSONSerializer implements JsonSerializer<Object>{

    @Override
    public JsonElement serialize(Object ref, Type arg1, JsonSerializationContext arg2) {
        return null;
    }
    
}
