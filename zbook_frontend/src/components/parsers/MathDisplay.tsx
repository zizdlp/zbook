import "katex/dist/katex.min.css";
import Latex from "react-latex-next";

interface MathDisplayProps {
  children: React.ReactNode;
}

const MathDisplay: React.FC<MathDisplayProps> = ({ children }) => (
  <Latex>{children?.toLocaleString() ?? ""}</Latex>
);

export default MathDisplay;
