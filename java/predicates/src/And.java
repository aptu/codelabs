public class And implements Predicate {
    private Predicate a;
    private Predicate b;

    public And(Predicate a, Predicate b) {
        this.a = a;
        this.b = b;
    }

    public boolean test(Instance i) {
        return this.a.test(i) && this.b.test(i);
    }
}
