import Card from "./card";
import { userService } from "services";

import FormData from "./form";

export default DashBoard;
function DashBoard() {
  const plateForm = ["LeetCode", "CodeChef", "InterviewBit", "HackerRank"];
  const userNames = userService.userValue?.accounts;
  return (
    <div>
      {!userNames.length ? <FormData></FormData> : null}
      <div className="grid grid-flow-col gap-3 p-3">
        {userNames.map((userName, index) => {
          if (userName !== "") {
            return (
              <Card
                plat={plateForm[index]}
                userName={userName}
                solved={10}
              ></Card>
            );
          }
        })}
      </div>
    </div>
  );
}
