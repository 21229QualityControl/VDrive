import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.io.PrintWriter;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.EnumMap;
import java.util.Map;

public class VDrive {
  ServerSocket serv;
  Socket conn;
  Thread thread;
  Map<Key, Boolean> pressed;
  boolean running;

  public VDrive() {
    pressed = new EnumMap<>(Key.class);
  }

  // Note: keybinds strings cant contain colon or semicolon
  public void start(int port, Map<Key, String> keybinds) throws IOException {
    serv = new ServerSocket(port);
    System.out.println("Waiting for connection on port " + port + "...");
    conn = serv.accept();

    // Generate keybind text
    StringBuilder sb = new StringBuilder();
    for (Map.Entry<Key, String> e : keybinds.entrySet()) {
      sb.append(e.getKey().name() + ": " + e.getValue() + ";");
    }
    PrintWriter out = new PrintWriter(conn.getOutputStream(), true);
    out.println(sb.toString());
    BufferedReader in = new BufferedReader(new InputStreamReader(conn.getInputStream()));

    // Start listening
    thread = new Thread(new Runnable() {
      public void run() {
        try {
          while (running) {
            String[] line = in.readLine().split(":");
            if (line == null) {
              break;
            }
            Key key = Key.fromString(line[0]);
            pressed.put(key, line[1].equals("true"));
            System.out.println("Pressed " + line[0]);
          }
        } catch (IOException e) {
          e.printStackTrace();
        }
      }
    });
    running = true;
    thread.setName("VDrive");
    thread.setPriority(Thread.NORM_PRIORITY);
    thread.start();
  }

  public void stop() throws IOException {
    running = false;
    conn.close();
    serv.close();
  }

  public boolean isDown(Key key) {
    return pressed.getOrDefault(key, false);
  }

  public enum Key {
    UP,
    DOWN,
    LEFT,
    RIGHT,
    Q,
    W,
    E,
    R,
    T,
    Y,
    U,
    I,
    O,
    P,
    A,
    S,
    D,
    F,
    G,
    H,
    J,
    K,
    L,
    Z,
    X,
    C,
    V,
    B,
    N,
    M,
    COMMA,
    PERIOD,
    SLASH,
    SEMICOLON,
    APOSTROPHE,
    LBRACKET,
    RBRACKET,
    MINUS,
    EQUAL,
    ONE,
    TWO,
    THREE,
    FOUR,
    FIVE,
    SIX,
    SEVEN,
    EIGHT,
    NINE,
    ZERO;

    public String string() {
      switch (this) {
        case UP: return "UP";
        case DOWN: return "DOWN";
        case LEFT: return "LEFT";
        case RIGHT: return "RIGHT";
        case Q: return "Q";
        case W: return "W";
        case E: return "E";
        case R: return "R";
        case T: return "T";
        case Y: return "Y";
        case U: return "U";
        case I: return "I";
        case O: return "O";
        case P: return "P";
        case A: return "A";
        case S: return "S";
        case D: return "D";
        case F: return "F";
        case G: return "G";
        case H: return "H";
        case J: return "J";
        case K: return "K";
        case L: return "L";
        case Z: return "Z";
        case X: return "X";
        case C: return "C";
        case V: return "V";
        case B: return "B";
        case N: return "N";
        case M: return "M";
        case COMMA: return ",";
        case PERIOD: return ".";
        case SLASH: return "/";
        case SEMICOLON: return ";";
        case APOSTROPHE: return "'";
        case LBRACKET: return "[";
        case RBRACKET: return "]";
        case MINUS: return "-";
        case EQUAL: return "=";
        case ONE: return "1";
        case TWO: return "2";
        case THREE: return "3";
        case FOUR: return "4";
        case FIVE: return "5";
        case SIX: return "6";
        case SEVEN: return "7";
        case EIGHT: return "8";
        case NINE: return "9";
        case ZERO: return "0";
        default: return "";
      }
    } 

    public static Key fromString(String val) {
      switch (val) {
        case "UP": return UP;
        case "DOWN": return DOWN;
        case "LEFT": return LEFT;
        case "RIGHT": return RIGHT;
        case "Q": return Q;
        case "W": return W;
        case "E": return E;
        case "R": return R;
        case "T": return T;
        case "Y": return Y;
        case "U": return U;
        case "I": return I;
        case "O": return O;
        case "P": return P;
        case "A": return A;
        case "S": return S;
        case "D": return D;
        case "F": return F;
        case "G": return G;
        case "H": return H;
        case "J": return J;
        case "K": return K;
        case "L": return L;
        case "Z": return Z;
        case "X": return X;
        case "C": return C;
        case "V": return V;
        case "B": return B;
        case "N": return N;
        case "M": return M;
        case ",": return COMMA;
        case ".": return PERIOD;
        case "/": return SLASH;
        case ";": return SEMICOLON;
        case "'": return APOSTROPHE;
        case "[": return LBRACKET;
        case "]": return RBRACKET;
        case "-": return MINUS;
        case "=": return EQUAL;
        case "1": return ONE;
        case "2": return TWO;
        case "3": return THREE;
        case "4": return FOUR;
        case "5": return FIVE;
        case "6": return SIX;
        case "7": return SEVEN;
        case "8": return EIGHT;
        case "9": return NINE;
        case "0": return ZERO;
        default: return null;
      }
    }
  }
}