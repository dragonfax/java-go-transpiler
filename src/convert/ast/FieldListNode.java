package convert.ast;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;

public class FieldListNode implements Node {

    FieldNode[] fields;

    @Override
    public String toString() {
        String[] strings = Arrays.stream(fields).map(Object::toString).toArray(String[]::new);
        return String.join("\n",strings);
    }

    public FieldListNode(FieldNode[] fields) {
        this.fields = fields;
    }

    public FieldListNode append(FieldListNode nextFieldList) {
        var list = new ArrayList<FieldNode>(fields.length + nextFieldList.fields.length);
        Collections.addAll(list, this.fields);
        Collections.addAll(list, nextFieldList.fields);
        return new FieldListNode(list.toArray<FieldNode>());
    }
    
}
