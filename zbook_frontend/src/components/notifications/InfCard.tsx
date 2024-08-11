import { useEffect, useRef } from "react";
import { ReactNode } from "react";
export default function InfCard({
  newLimit,
  isLast,
  children
}: {
  newLimit: any;
  isLast: boolean;
  children: ReactNode;
}) {
  /**
   * Select the Card component with useRef
   */
  // const cardRef = useRef();
  const cardRef = useRef<HTMLDivElement | null>(null);

  /**
   * Implement Intersection Observer to check if the last Card in the array is visible on the screen, then set a new limit
   */
  useEffect(() => {
    if (!cardRef?.current) return;
    const observer = new IntersectionObserver(([entry]) => {
      if (isLast && entry.isIntersecting) {
        newLimit();
        observer.unobserve(entry.target);
      }
    });
    observer.observe(cardRef.current);
  }, [isLast, newLimit]);

  return (
    <div ref={cardRef}>
      {children}
    </div>
  );
}
