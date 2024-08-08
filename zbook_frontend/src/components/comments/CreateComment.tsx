"use client";
import { ThemeColor } from "../TableOfContent";
import CreateCommentForm from "./CreateCommentForm";
interface CommentFormProps {
  markdownID: number;
  parentID: number;
  username: string;
  theme_color: ThemeColor;
}
export default function CreateComment(props: CommentFormProps) {
  return (
    <div className="flex flex-row space-x-4 my-2 py-2 items-center">
      <CreateCommentForm
        markdownID={props.markdownID}
        parentID={props.parentID}
        username={props.username}
        theme_color={props.theme_color}
      />
    </div>
  );
}
