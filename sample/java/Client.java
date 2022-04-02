import uz.yt.eimzo.dsv.server.plugin.pkcs7.v1.Pkcs7Service;
import uz.yt.eimzo.dsv.server.plugin.pkcs7.v1.Pkcs7;

public class Client {

    public static void main(String[] args){
        Pkcs7Service service = new Pkcs7Service();
        Pkcs7 port = service.getPkcs7Port();

        String pkcs7B64 = "MIIikAYJKoZI.................";

        String result = port.verifyPkcs7(pkcs7B64);
        System.out.println(result);
    }   
}