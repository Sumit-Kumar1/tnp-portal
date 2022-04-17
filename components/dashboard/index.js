import Card from "./card";
export default DashBoard;

function DashBoard() {
  return (
    <div>
      <div className="flex flex-row gap-5 my-4 mx-auto">
        <Card
          plat="LeetCode"
          userName="Dark_Blood"
          solved={Math.floor(Math.random() * 100)}
        />
        <Card
          plat="LeetCode"
          userName="Dark_Blood"
          solved={Math.floor(Math.random() * 100)}
        />
        <Card
          plat="LeetCode"
          userName="Dark_Blood"
          solved={Math.floor(Math.random() * 100)}
        />
      </div>
    </div>
  );
}
