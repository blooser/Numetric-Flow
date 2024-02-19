'use client'

import Image from "next/image";
import { useState } from "react";
import Button from "./components/Button";
import NumberInput from "./components/NumberInput";
import Result from "./components/Result";
import Separator from "./components/Separator";
import { useNumber } from "./context/numbers";


export default function Home() {
  const [result, setResult] = useState(null)
  const {number, getIndex} = useNumber()

  const handleOnCLick = () => {
    getIndex(number).then(response => {
      if (response !== false) {
        setResult(response.index)
      }
    })
  }

  return (
    <main className="flex min-h-screen flex-col space-y-12 items-center justify-center p-24">
      <Image src="/logo.webp" width="256" height="124" className="w-15 h-15  rounded-full"/>
      <Separator />
      <div className="flex flex-row space-x-4">
          <NumberInput />
          <Button onClick={handleOnCLick}/>
        </div>
        {result !== null &&
          <Result number={result}/>
        }
    </main>
  );
}
