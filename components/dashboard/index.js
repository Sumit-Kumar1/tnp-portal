import Card from "./card";
import { userService } from "services";
export default DashBoard;
import FormData from "./form";
import {useState, useEffect} from 'react';

function DashBoard() {
  console.log(userService.user)
  return (
    <div>
      {!userService.user.accounts ? <FormData></FormData> : null}
      <div className="grid grid-flow-col gap-3 p-3">
        <Card
          plat="leetCode"
          userName="Dark_Blood"
          solved={Math.floor(Math.random() * 100)}
        />
        <Card
          plat="interviewBit"
          userName="Dark_Blood"
          solved={Math.floor(Math.random() * 100)}
        />
        <Card
          plat="codechef"
          userName="Dark_Blood"
          solved={Math.floor(Math.random() * 100)}
        />
      </div>
    </div>
  );
}
