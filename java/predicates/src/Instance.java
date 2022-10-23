import java.util.List;

public class Instance {
    String type;
    String os;
    List<String> tags;
    int launchTime;

    public Instance(String type, String os, List<String> tags, int launchTime) {
        this.type = type;
        this.os = os;
        this.tags = tags;
        this.launchTime = launchTime;
    }

    public String getOs() {
        return os;
    }

    public void setOs(String os) {
        this.os = os;
    }

    public List<String> getTags() {
        return tags;
    }

    public void setTags(List<String> tags) {
        this.tags = tags;
    }

    public int getLaunchTime() {
        return launchTime;
    }

    public void setLaunchTime(int launchTime) {
        this.launchTime = launchTime;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }
}
