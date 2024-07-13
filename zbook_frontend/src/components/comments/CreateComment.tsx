"use client";
import CreateCommentForm from "./CreateCommentForm";
interface CommentFormProps {
  markdownID: number;
  parentID: number;
  username: string;
}
export default function CreateComment(props: CommentFormProps) {
  return (
    <div className="flex flex-row space-x-4 my-2 py-2 items-center">
      <CreateCommentForm
        markdownID={props.markdownID}
        parentID={props.parentID}
        username={props.username}
      />
    </div>
  );
}
