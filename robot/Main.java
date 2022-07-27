import java.io.IOException;
import java.util.EnumMap;
import java.util.Map;

public class Main {
  public static void main(String[] args) throws IOException {
    VDrive robot = new VDrive();
    Map<VDrive.Key, String> keybinds = new EnumMap<>(VDrive.Key.class);
    keybinds.put(VDrive.Key.W, "Forwards");
    keybinds.put(VDrive.Key.A, "Left");
    keybinds.put(VDrive.Key.S, "Backwards");
    keybinds.put(VDrive.Key.D, "Right");
    robot.start(8080, keybinds);

    while (!robot.isDown(VDrive.Key.W)) {
      // Do nothing
    }

    robot.stop();
  }
}
