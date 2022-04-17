import { userService } from "services";
import { Link } from "components";
import DashBoard from "components/dashboard";

export default Home;

function Home() {
  return (
    <div className="p-4">
      <div className="container">
        <h1 className="text-xl font-bold">
          Hi {userService.userValue?.firstName}! ({userService.userValue?.ErNo})
        </h1>
        <DashBoard></DashBoard>
        <p>
        </p>
      </div>
    </div>
  );
}
