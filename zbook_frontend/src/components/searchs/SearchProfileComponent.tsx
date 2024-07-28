import { Link } from "@/navigation";
import { ListUserInfo } from "@/types/model";
import AvatarImageClient from "../AvatarImageClient";
import SearchItemWrapper from "./SearchItemWrapper";
interface ProfileProps {
  ListUserInfo: ListUserInfo;
}

export default function SearchProfileComponent(props: ProfileProps) {
  return (
    <SearchItemWrapper>
      <Link
        href={"/workspace/" + props.ListUserInfo.username}
        className="flex items-center justify-start w-full px-2"
      >
        <AvatarImageClient
          username={props.ListUserInfo.username}
          className="w-12 h-12 mx-4 rounded-full shadow-lg"
        />
        <div className="flex flex-col py-2">
          <strong className=" text-sm font-medium  flex space-x-2 items-center justify-start">
            <p>{props.ListUserInfo.username}</p>
          </strong>
          <span className=" text-sm font-medium ">
            {props.ListUserInfo.email}
          </span>
        </div>
      </Link>
    </SearchItemWrapper>
  );
}
