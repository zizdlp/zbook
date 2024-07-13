export default function NavBarIcon({
  Icon,
  onClick,
  mounted,
}: {
  Icon: any;
  onClick: () => void;
  mounted: boolean;
}) {
  if (!mounted) {
    return (
      <div className="block w-6 h-6 bg-slate-200 dark:bg-slate-500  animate-pulse rounded-full" />
    );
  }
  return (
    <Icon
      onClick={onClick}
      className="block w-6 h-6  hover:text-sky-600 dark:hover:text-sky-400 cursor-pointer"
    />
  );
}
