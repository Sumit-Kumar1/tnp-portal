import Image from "next/image";
import leetCode from "../../public/leetB.svg";
// import codechef from "../..s/public/Codecheflogo.svg.png";
import interviewBit from "../../public/interviewBit.png";
import hackkerank from "../../public/hackerrank.svg";
import codeforces from "../../public/codeforces.svg";
import spoj from "../../public/SPOJ.svg";

function Card(props) {
  const plateform =
    props.plat === "leetcode"
      ? leetCode
      : props.plat === "codeforces"
      ? codeforces
      : props.plat === "interviewbit"
      ? interviewBit
      : props.plat === "spoj"
      ? spoj
      : hackkerank;
  return (
    <div className="grid grid-flow-row border-2 rounded-xl p-2 text-center md:grid-flow-col">
      <Image
        src={plateform}
        alt={`${props.plat} image`}
        layout="responsive"
        objectFit="cover"
        // height={100}
        // width={100}
        className=""
      ></Image>{" "}
      <div className="">
        <h1 className="font-bold text-xl py-2">{props.plat}</h1>
        <h2 className="text-lg text-indigo-500">{props.userName}</h2>
        <p><span className="font-light">{`( ${props.solved} )`}</span></p>
      </div>
    </div>
  );
}

export default Card;
