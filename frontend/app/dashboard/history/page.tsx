"use client";

import CustomButton from "@/app/components/button";
import { Dropdown } from "@/app/components/dropdown";
import { InputWithLabel } from "@/app/components/input";

const History = () => {
  return (
    <div className="shadow-[0px_0px_32px_0px_rgba(204,204,204,0.25)] rounded-2xl px-10 py-10">
      <form
        action=""
        className="flex gap-5 border-b border-[#EBEBEB] pb-5 mb-5"
      >
        <InputWithLabel className="rounded-2xl" label="Number" />
        <InputWithLabel className="rounded-2xl" label="Name" />
        <InputWithLabel className="rounded-2xl" label="Staff" />
        <InputWithLabel className="rounded-2xl" label="Engineer" />
        <InputWithLabel className="rounded-2xl" label="Status" />
        <InputWithLabel className="rounded-2xl" label="Priority" />
        <InputWithLabel className="rounded-2xl" label="Title" />
      </form>
      <div className="flex justify-between">
        <Dropdown />
        <div className="gap-5 flex">
          <CustomButton>Previous</CustomButton>
          <CustomButton>Next</CustomButton>
        </div>
      </div>
    </div>
  );
};

export default History;
