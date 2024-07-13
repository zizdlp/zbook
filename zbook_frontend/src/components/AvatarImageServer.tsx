import Image from "next/image";
import LoadingElement from "./loadings/LoadingElement";
import { logger } from "@/utils/logger";
import { FetchError } from "@/fetchs/util";
import { getUserAvatarServer } from "@/fetchs/server_without_auth";

export default async function AvatarImageServer({
  username,
  className,
}: {
  username: string;
  className: string;
}) {
  try {
    const data = await getUserAvatarServer({ username });
    if (data.avatar) {
      const avatar = `data:image/png;base64,${data.avatar}`;
      return (
        <Image
          src={avatar}
          className={className}
          alt="avatar"
          width={50}
          height={50}
        />
      );
    } else {
      throw new FetchError("server fetch avatar failed", 404);
    }
  } catch (error) {
    logger.error(`get avatar from server side failed:${error}`);
    return <LoadingElement className={`${className} animate-pulse`} />;
  }
}
