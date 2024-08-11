"use client";

import DialogComponent from "../DialogComponent";
import { OperationContext } from "@/providers/OperationProvider";
import React, { useContext, useEffect, useState } from "react";
import CreateCommentForm from "@/components/comments/CreateCommentForm";
import { useSession } from "next-auth/react";
import { ThemeColor } from "../TableOfContent";
export default function CreateCommentDialog() {
  const [username, setUsername] = useState("");
  const { data, status } = useSession();
  useEffect(() => {
    if (data?.username) {
      setUsername(data.username);
    }
  }, [data]);

  const {
    createCommentOpen,
    setCreateCommentOpen,
    operationMarkdownID,
    operationParentID,
  } = useContext(OperationContext);

  return (
    <DialogComponent
      showDialog={createCommentOpen}
      setShowDialog={setCreateCommentOpen}
    >
      <div className="my-4 items-center justify-center mx-4">
        <CreateCommentForm
          markdownID={operationMarkdownID}
          parentID={operationParentID}
          username={username}
          theme_color={ThemeColor.Sky}
        />
      </div>
    </DialogComponent>
  );
}
