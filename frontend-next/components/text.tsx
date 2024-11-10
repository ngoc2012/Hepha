'use client'

import { editText } from '@/SSR/text';
import React, { useState, useRef, useEffect } from 'react';

interface Props {
  table : string
  id: string
  field: string
  value: string
  format: string
}

// format: text-4xl
const Text: React.FC<Props> = ({ table, id, field, value, format }) => {
  const [isEditing, setIsEditing] = useState(false);
  const [text, setText] = useState(value);
  const inputRef = useRef<HTMLInputElement>(null);
  // const format = "text-4xl"; // store the class name in a variable

  useEffect(() => {
    if (isEditing) {
      inputRef.current?.focus();
    }
  }, [isEditing]);

  const handleClick = () => {
    setIsEditing(true);
  };

  // `http://${process.env.GIN_ADDR}/edit/${id}`

  const handleBlur = async () => {
    setIsEditing(false);
    await editText(table, id, field, text);
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setText(e.target.value);
  };

  return (
    <div className="inline-block">
      {isEditing ? (
        <input
          ref={inputRef}
          type="text"
          value={text}
          onChange={handleChange}
          onBlur={handleBlur}
          className={`${format} px-2 py-1 border border-gray-300 rounded focus:outline-none focus:border-blue-500`}
        />
      ) : (
        <span onClick={handleClick} className={`${format} cursor-text`}>
          {text}
        </span>
      )}
    </div>
  );
};

export default Text ;
