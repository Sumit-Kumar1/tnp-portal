import Card from "./card";
import { userService } from "services";
import FormData from "./form";
import { LeetCodeData } from "helpers/fetchLeetCode";

export default DashBoard;
function DashBoard() {
  const plateForm = ["LeetCode", "CodeChef", "InterviewBit", "HackerRank"];
  const userNames = userService.userValue?.accounts;
  return (
    <div>
      {!userService.userValue?.accounts.length ? (
        <FormData></FormData>
      ) : (
        <div className="grid grid-flow-col gap-3 p-3">
          {userNames.map((userName, index) => {
            if (userName !== "") {
              const {data, error} = LeetCodeData(userName);
              if(error) return <h1>Error occured during Fetching</h1>;
              if(!data) return <p>Loading Data...</p>

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
      )}
    </div>
  );
}
