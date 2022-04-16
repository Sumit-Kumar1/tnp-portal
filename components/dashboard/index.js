import Card from "./card";
import { userService } from "services";
export default DashBoard;


function DashBoard() {
    if(userService.user)
  return (
    <div className="flex flex-row gap-5 my-4 mx-auto">
      <Card plat="LeetCode" userName="Dark_Blood" solved={Math.floor(Math.random()*100)}/>
      <Card plat="LeetCode" userName="Dark_Blood" />
      <Card plat="LeetCode" userName="Dark_Blood" />
    </div>
  );
}

