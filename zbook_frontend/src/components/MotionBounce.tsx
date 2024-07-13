"use client";
import { motion, Variants } from "framer-motion";

const cardVariants: Variants = {
  offxscreen: {
    y: 100,
    opacity: 0,
  },
  offyscreen: {
    y: 100,
    opacity: 0,
  },
  onscreen: {
    x: 0,
    y: 0,
    opacity: 1,
    transition: {
      type: "spring",
      bounce: 0.2,
      duration: 1.5,
    },
  },
};

export default function MotionBounce({
  children,
  direction,
}: {
  children: React.ReactNode;
  direction: string;
}) {
  return (
    <motion.div
      initial={`${direction == "x" ? "offxscreen" : "offyscreen"}`}
      whileInView="onscreen"
      viewport={{ once: true, amount: 0.1 }}
    >
      <motion.div variants={cardVariants}>{children}</motion.div>
    </motion.div>
  );
}
