import com.rabbitmq.client.Channel;
import com.rabbitmq.client.Connection;
import com.rabbitmq.client.ConnectionFactory;

public class Sol {

    public static void main(String[] args) throws Exception {
        ConnectionFactory factory = new ConnectionFactory();
        User user=new User();
        user.setUsername("David");
        System.out.println(user.getUsername());
    }
}