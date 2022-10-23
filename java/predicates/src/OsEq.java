

public class OsEq implements Predicate {
    private String os;
    public OsEq(String os) {
        this.os = os;
    }

    @Override
    public boolean test(Instance i) {
        return i.getOs().equals(os);
    }
}
