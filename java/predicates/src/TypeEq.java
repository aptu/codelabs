
public class TypeEq implements Predicate {
    String type;

    public TypeEq(String type) {
        this.type = type;
    }
    @Override
    public boolean test(Instance i) {
        return i.getType().equals(this.type);
    }
}
