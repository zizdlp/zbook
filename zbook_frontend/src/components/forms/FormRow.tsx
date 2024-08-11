export default function FormRow({
  children,
  error,
  show_error,
}: {
  children: React.ReactNode;
  error: string | undefined;
  show_error: boolean | undefined;
}) {
  return (
    <div className="mb-2">
      <div className="flex justify-center items-center relative">
        {children}
      </div>

      <div className="flex items-center justify-begin h-4 pt-1 text-xs text-pink-600">
        {error && show_error && (error as string)}
      </div>
    </div>
  );
}
