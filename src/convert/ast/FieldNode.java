package convert.ast;

public class FieldNode implements Node{

    String name;
    String type;

    @Override
    public String toString() {
        return String.format("%s %s", name, type);
    }

    public FieldNode(String name, String type) {
        this.name = name;
        this.type = type;
    }
    
}
