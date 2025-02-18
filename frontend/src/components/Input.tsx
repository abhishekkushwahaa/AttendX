// eslint-disable-next-line @typescript-eslint/no-explicit-any
const Input = ({ type, placeholder, value, onChange }: any) => (
  <input
    type={type}
    placeholder={placeholder}
    value={value}
    onChange={onChange}
  />
);

export default Input;
