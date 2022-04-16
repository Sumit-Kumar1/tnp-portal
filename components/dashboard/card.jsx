import Image from "next/image";
import leetCodeImage from "../../public/leetB.svg";
function Card(props){
    return (
        <div className="border-2 p-3 rounded-lg w-full text-center border-green-500">
            <Image src={leetCodeImage} alt="LeetCode-image" objectFit="cover"></Image>
            <h1>{props.plat}</h1>
            <h2>{props.userName}</h2>
            <p>Problems Solved: {props.solved}</p>
        </div>
    );
}

export default Card;