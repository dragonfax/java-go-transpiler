package convert.ast;

public class FieldNode implements Node{

    public String name;
    public String type;

    @Override
    public String toString() {
        return String.format("%s %s", name, type);
    }

    public FieldNode(String name, String type) {
        this.name = name;
        this.type = type;
    }
    
}
