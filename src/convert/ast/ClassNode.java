package convert.ast;

public class ClassNode implements Node {

    String name;
    FieldListNode fields;

    @Override
    public String toString() {
        return String.format("type %s struct {\n%s\n}\n", name, fields.toString());
    }

    public ClassNode(String name, FieldListNode fields) {
        this.name = name;
        this.fields = fields;
    }
}
