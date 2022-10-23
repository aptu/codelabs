import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.stream.Collectors;

public class Predicates {



    public boolean matchType(Instance instance, String type){
        return instance.getType().equals(type);
    }

    public boolean matchOs(Instance instance, String os){
        return instance.getOs().equals(os);
    }




    public static void main(String[] args) {

        List<Instance> instances = Arrays.asList(
                new Instance("t3.nano", "Linux", Arrays.asList("production", "worker"), 100),
                new Instance("t3.micro", "Linux", Arrays.asList("gamma", "service"), 101),
                new Instance("m5.nano", "MacOS", Arrays.asList("production", "worker"), 102),
                new Instance("c5.nano", "Linux", Arrays.asList("production", "worker"), 103)
        );

        // write code which enables us to search for instances that match a certain criteria,
        // e.g. 'find all t3.nano instances'
        // or 'find all Linux instances'

        Predicates p = new Predicates();

        Predicate typeEq = new TypeEq("t3.nano");
        Predicate osEq = new OsEq("Linux");
        for (Instance i: instances) {
            //System.out.println((new TypeEq("t3.nano").test(i) && new OsEq("MacOS").test( i)));
            System.out.println(new And(typeEq, osEq).test(i));
            System.out.println(typeEq.and(osEq).test(i));

        }

        System.out.println(instances
                .stream()
                .filter(typeEq.and(osEq))
                .collect(Collectors.toList())
        );
    }

}

