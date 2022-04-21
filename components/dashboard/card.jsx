import Image from "next/image";
import leetCode from "../../public/leetB.svg";
import codechef from "../../public/Codecheflogo.svg.png";
import interviewBit from "../../public/interviewBit.png";
import hackkerank from "../../public/hackerrank.svg";

function Card(props) {
  const plateform =
    props.plat === "LeetCode"
      ? leetCode
      : props.plat === "CodeChef"
      ? codechef
      : props.plat === "InterviewBit"
      ? interviewBit
      : hackkerank;
  return (
    <div className="flex flex-col border border-indigo-500 rounded-xl p-3 text-center">
      <Image
        src={plateform}
        alt={`${props.plat} image`}
        layout="responsive"
        objectFit="cover"
        height={100}
        width={100}
      ></Image>{" "}
      <div>
        <h1 className="font-bold text-xl py-2">{props.plat}</h1>
        <h2>{props.userName}</h2>
        <p>Problems Solved: {props.solved}</p>
      </div>
    </div>
  );
}

export default Card;
