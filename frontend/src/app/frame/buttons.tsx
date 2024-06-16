import styles from "src/styles";

type ButtonProps = {
  text: string;
  callback?: () => void;
};

const DefaultButton = (props: ButtonProps) => {
  return (
    <button className={styles.buttons.collect}>
      <div onClick={props.callback}>
        <span>{props.text}</span>
      </div>
    </button>
  );
};

const GreenButton = (props: ButtonProps) => {
  return (
    <button className={styles.buttons.green}>
      <div onClick={props.callback}>
        <span>{props.text}</span>
      </div>
    </button>
  );
};

export { DefaultButton as CollectButton, GreenButton };
