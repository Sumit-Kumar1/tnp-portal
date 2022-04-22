import useSWR from "swr";

const urls = "https://leetcode-stats-api.herokuapp.com/";
const fetcher = (url) => fetch(url).then((res) => res.json());
export const LeetCodeData = (userName) => {
  if (!userName) {
    throw new Error("User Name is required");
  }
  const url = urls + userName;
  const { data, error } = useSWR(url, fetcher);
  return { data, error };
};
