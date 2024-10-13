public class Main {
    static {
        System.loadLibrary("greetings");
    }

    private native String GreetPerson(String name);

    public static void main(String[] args) {
        String greeting = new Main().GreetPerson("Guy");
        System.out.println(greeting); // Should output "Hello, Guy!"
    }
}
