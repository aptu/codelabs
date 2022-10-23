import java.util.LinkedList;

public interface Predicate extends java.util.function.Predicate {
    boolean test(Instance i);
    default boolean test(Object o) {
        return this.test((Instance)o);
    }

    default Predicate and(Predicate p) {
        return new And(this, p);
    }

}
