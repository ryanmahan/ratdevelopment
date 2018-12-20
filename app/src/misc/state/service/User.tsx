export class UserService {

    static logout() {
      localStorage.removeItem("state");
      localStorage.removeItem('email');
    }

}
