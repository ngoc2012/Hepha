'use client'

// components/Box.tsx
import React, {useState} from 'react';
import {exeBox} from '../SSR/box';

interface Props {
  value: string;
  onChange: (event: React.ChangeEvent<HTMLTextAreaElement>) => void;
  onKeyDown: (event: React.KeyboardEvent<HTMLTextAreaElement>) => void;
  placeholder?: string;
  rows?: number;
  cols?: number;
}

const TextArea: React.FC<Props> = ({
  value,
  onChange,
  onKeyDown,
  placeholder = '',
  rows = 4,
  cols = 50,
}) => {
  return (
    <textarea
      className="border border-gray-400 rounded font-sans"
      value={value}
      onChange={onChange}
      onKeyDown={onKeyDown}
      placeholder={placeholder}
      rows={rows}
      cols={cols}
    />
  );
};

const Box: React.FC = () => {
  const [text, setText] = useState('');
  const [output, setOutput] = useState('');
  const [lang, setLang] = useState('Python');
  const [ver, setVer] = useState('3');
  const langs = [
    {'lang': 'Python', 'ver': ['2', '3']},
    {'lang': 'C', 'ver': []}
  ]

  const handleButtonClick = async () => {
    const res = await exeBox("Code", text, lang, ver);
    setOutput(res);
  };

  const handleKeyDown = (event: React.KeyboardEvent<HTMLTextAreaElement>) => {
    // console.log('Key pressed:', event.key);
    if (event.key === 'Tab') {
      event.preventDefault();
      // console.log('Tab pressed');

      // Get the current cursor position
      const cursorPosition = event.currentTarget.selectionStart;

      // Create a new text with the tab character inserted at the cursor position
      const newText = [
        text.slice(0, cursorPosition),
        '\t',
        text.slice(cursorPosition),
      ].join('');

      // Update the text state with the new text
      setText(newText);

      // Set the cursor position to right after the inserted tab character
      event.currentTarget.selectionStart = cursorPosition + 1;
      event.currentTarget.selectionEnd = cursorPosition + 1;
    }
  };

  const handleLangChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedLang = e.target.value;
    setLang(selectedLang);
    const selectedLangObj = langs.find(l => l.lang === selectedLang);
    if (selectedLangObj && selectedLangObj.ver.length > 0) {
      setVer(selectedLangObj.ver[0]);
    } else {
      setVer('');
    }
  };

  const handleVerChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setVer(e.target.value);
  };

  const selectedLangObj = langs.find(l => l.lang === lang);
  const versions = selectedLangObj ? selectedLangObj.ver : [];

  return (
    <div className="border-2 border-gray-400 rounded">
      <div className="flex space-x-4">
        <div>
          <label htmlFor="lang">Language:</label>
          <select id="lang" value={lang} onChange={handleLangChange}>
            {langs.map(l => (
              <option key={l.lang} value={l.lang}>
                {l.lang}
              </option>
            ))}
          </select>
        </div>
        <div>
          <label htmlFor="ver">Version:</label>
          <select id="ver" value={ver} onChange={handleVerChange} disabled={versions.length === 0}>
            {versions.map(v => (
              <option key={v} value={v}>
                {v}
              </option>
            ))}
          </select>
        </div>
      </div>
      <button
        className="bg-blue-500 hover:bg-blue-700 text-black font-mono font-bold py-2 px-4 rounded"
        onClick={handleButtonClick}
      >
        Run
      </button>
      <TextArea value={text} onChange={(event) => setText(event.target.value)} onKeyDown={handleKeyDown}/>
      <div id="output">{output}</div>
    </div>
  );
};

export default Box;
