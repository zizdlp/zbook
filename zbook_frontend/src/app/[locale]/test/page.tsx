"use client";
// import Editor from "@/editor/index";
import Editor from "@/editor/stories/index";
export default function Home() {
  return (
    <div className="h-full w-full mx-auto max-w-7xl">
      {/* <Editor defaultValue="Hello world!" /> */}
      <Editor
        defaultValue="Hello world!"
        readOnly={false}
        onSave="save"
        onCancel="cancel"
      />
      <h1 className="text-3xl font-semibold">nihao</h1>
    </div>
  );
}
