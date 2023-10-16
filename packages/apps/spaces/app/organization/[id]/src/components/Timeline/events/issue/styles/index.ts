export const CustomTicketTearStyle = {
  '&:before': {
    content: "''",
    width: '11px',
    height: '5px',
    border: '1px solid',
    borderColor: 'gray.200',
    left: '-7px',
    top: '-1px',
    position: 'absolute',
    background: 'gray.25',
    borderTop: 'none',
    borderBottomLeftRadius: '10px',
    borderBottomRightRadius: '10px',
    zIndex: 1,
  },
  '&:after': {
    content: "''",
    width: '11px',
    height: '5px',
    border: '1px solid',
    borderColor: 'gray.200',
    left: '-7px',
    bottom: '-1px',
    position: 'absolute',
    background: 'gray.25',
    borderBottomColor: 'gray.25',
    borderTopLeftRadius: '10px',
    borderTopRightRadius: '10px',
    zIndex: 1,
    boxShadow: 'inset 0px 1px 2px 0px rgba(16, 24, 40, 0.05)',
  },
};
