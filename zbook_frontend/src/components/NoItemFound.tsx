import { FaBox } from "react-icons/fa";

export default function NoItemFound({ title }: { title: string }) {
    return (<div className="h-full flex w-full items-center justify-center ">
        <div className="flex flex-col items-center space-y-2">
            <FaBox className="h-12 w-12 dark:text-slate-700/50 text-slate-400/50" />
            <p className="py-2 font-semibold text-lg dark:text-slate-600/50 text-slate-500/50">
                {title}
            </p>
        </div>
    </div>)
}