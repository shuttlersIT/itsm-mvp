import { Button } from "@/components/ui/button";
const CustomButton = ({
  children,
  className,
}: {
  children: React.ReactNode;
  className?: string;
}) => {
  return (
    <Button
      className={`${
        className?.includes("bg")
          ? `${className}`
          : `bg-white hover:bg-[#F4F4F4] border-[#EBEBEB] ${className}`
      } text-xs md:text-[0.95rem] text-[#060606] hover:border-transparent gap-2 border rounded-full md:rounded-3xl shadow-[0px_0px_32px_0px_rgba(204,204,204,0.25)]`}
    >
      {children}
    </Button>
  );
};

export default CustomButton;
