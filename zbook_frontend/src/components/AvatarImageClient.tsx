"use client";
import Image from "next/image";
import { useEffect, useState, useMemo } from "react";
import LoadingElement from "./loadings/LoadingElement";
import { logger } from "@/utils/logger";
import { getUserAvatarServer } from "@/fetchs/server_without_auth";

export default function AvatarImageClient({
  username,
  className,
}: {
  username: string;
  className: string;
}) {
  const [userImage, setUserImage] = useState<string>();

  useEffect(() => {
    getUserAvatarServer({ username })
      .then((data: any) => {
        setUserImage(data.avatar);
      })
      .catch((error) => {
        logger.error(`Failed to fetch user avatar:${error}`);
      });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const avatar = useMemo(() => {
    return userImage ? `data:image/png;base64,${userImage}` : null;
  }, [userImage]);

  if (avatar) {
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
    return <LoadingElement className={`${className} animate-pulse`} />;
  }
}
