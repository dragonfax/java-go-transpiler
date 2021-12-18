package convert.ast;

import java.util.Arrays;

public class FieldListNode implements Node {

    FieldNode[] fields;

    @Override
    public String toString() {
        String[] strings = Arrays.stream(fields).map(Object::toString).toArray(String[]::new);
        return String.join("\n",strings);
    }
    
}
