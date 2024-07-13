"use client";
import { signIn } from "next-auth/react";
import { useEffect } from "react";

export default function CallSignIn({ oauthType }: { oauthType: string }) {
  useEffect(() => {
    signIn(oauthType);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
  return <></>;
}
