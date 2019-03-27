// @flow
import Box from './box'
import ClickableBox from './clickable-box'
import Icon, {castPlatformStyles, type IconType} from './icon'
import * as React from 'react'
import Text from './text'
import {
  type StylesCrossPlatform,
  borderRadius,
  collapseStyles,
  globalColors,
  globalStyles,
  globalMargins,
  isMobile,
  platformStyles,
  styleSheetCreate,
} from '../styles'

export type Props = {|
  children?: React.Node,
  onClick?: null | ((event: SyntheticEvent<>) => void),
  onMouseEnter?: Function,
  onMouseLeave?: Function,
  icon?: IconType,
  label?: string,
  style?: StylesCrossPlatform,
  labelContainerStyle?: StylesCrossPlatform,
  labelStyle?: StylesCrossPlatform,
  type:
    | 'Primary'
    | 'PrimaryPrivate'
    | 'Secondary'
    | 'Danger'
    | 'Wallet'
    | 'PrimaryGreen'
    | 'PrimaryGreenActive'
    | 'PrimaryColoredBackground'
    | 'SecondaryColoredBackground',
  disabled?: boolean,
  waiting?: boolean,
  small?: boolean,
  fullWidth?: boolean,
  backgroundMode?: 'Normal' | 'Terminal' | 'Red' | 'Green' | 'Blue' | 'Black' | 'Purple',
  className?: string,
|}

const Progress = ({small, white}) => (
  <Box style={styles.progressContainer}>
    <Icon
      style={castPlatformStyles(small ? styles.progressSmall : styles.progressNormal)}
      type={white ? 'icon-progress-white-animated' : 'icon-progress-grey-animated'}
    />
  </Box>
)

class Button extends React.Component<Props> {
  render() {
    const backgroundModeName = this.props.backgroundMode
      ? {
          Black: 'Black',
          Blue: 'Blue',
          Green: 'Green',
          Normal: '',
          Purple: 'Purple',
          Red: 'Red',
          Terminal: 'OnTerminal',
        }[this.props.backgroundMode]
      : ''

    let containerStyle = containerStyles[this.props.type + backgroundModeName]
    let labelStyle = labelStyles[this.props.type + 'Label' + backgroundModeName]

    if (this.props.fullWidth) {
      containerStyle = collapseStyles([containerStyle, styles.fullWidth])
    }

    if (this.props.small) {
      containerStyle = collapseStyles([containerStyle, styles.small])
    }

    if (this.props.icon) {
      containerStyle = collapseStyles([containerStyle, styles.icon])
    }

    if (this.props.disabled || this.props.waiting) {
      containerStyle = collapseStyles([containerStyle, styles.opacity30])
    }

    if (this.props.waiting) {
      labelStyle = collapseStyles([labelStyle, styles.opacity0])
    }

    containerStyle = collapseStyles([containerStyle, this.props.style])

    const onClick = (!this.props.disabled && !this.props.waiting && this.props.onClick) || null
    const whiteSpinner = !(
      this.props.type === 'PrimaryGreenActive' ||
      this.props.type === 'Secondary' ||
      this.props.type === 'PrimaryColoredBackground'
    )

    return (
      <ClickableBox
        style={containerStyle}
        onClick={onClick}
        onMouseEnter={this.props.onMouseEnter}
        onMouseLeave={this.props.onMouseLeave}
      >
        <Box
          style={collapseStyles([
            globalStyles.flexBoxRow,
            globalStyles.flexBoxCenter,
            styles.labelContainer,
            this.props.labelContainerStyle,
          ])}
        >
          {!this.props.waiting && this.props.children}
          {!!this.props.icon && (
            <Icon
              color={
                labelStyle.color === globalColors.black && !!this.props.label
                  ? globalColors.black_50
                  : labelStyle.color
              }
              sizeType={isMobile || !!this.props.label ? 'Small' : 'Default'}
              style={collapseStyles([!!this.props.label && styles.iconWithLabel])}
              type={this.props.icon}
            />
          )}
          {!!this.props.label && (
            <Text
              type={this.props.small ? 'BodySmallSemibold' : 'BodySemibold'}
              style={collapseStyles([labelStyle, this.props.labelStyle])}
            >
              {this.props.label}
            </Text>
          )}
          {!!this.props.waiting && <Progress small={this.props.small} white={whiteSpinner} />}
        </Box>
      </ClickableBox>
    )
  }
}

const smallHeight = isMobile ? 28 : 24
const regularHeight = isMobile ? 32 : 28
const fullWidthHeight = isMobile ? 48 : 40

const common = platformStyles({
  common: {
    ...globalStyles.flexBoxColumn,
    alignItems: 'center',
    alignSelf: 'center',
    borderRadius,
    height: regularHeight,
    justifyContent: 'center',
  },
  isElectron: {
    display: 'inline-block',
    lineHeight: 'inherit',
    paddingLeft: globalMargins.medium,
    paddingRight: globalMargins.medium,
  },
  isMobile: {
    paddingLeft: globalMargins.small,
    paddingRight: globalMargins.small,
  },
})

const commonLabel = platformStyles({
  common: {
    color: globalColors.white,
    textAlign: 'center',
  },
  isElectron: {whiteSpace: 'pre'},
  isMobile: {
    position: 'relative',
    top: 2,
  },
})

const styles = styleSheetCreate({
  fullWidth: {
    alignSelf: undefined,
    flexGrow: 1,
    height: fullWidthHeight,
    width: undefined,
  },
  icon: {
    paddingLeft: globalMargins.xsmall,
    paddingRight: globalMargins.xsmall,
  },
  iconWithLabel: {
    alignSelf: 'center',
    marginRight: globalMargins.tiny,
  },
  labelContainer: {height: '100%', position: 'relative'},
  opacity0: {opacity: 0},
  opacity30: {opacity: 0.3},
  progressContainer: {...globalStyles.fillAbsolute, ...globalStyles.flexBoxCenter},
  progressNormal: {height: isMobile ? 28 : 20},
  progressSmall: {height: isMobile ? 24 : 16},
  small: {
    borderRadius,
    height: smallHeight,
    paddingLeft: globalMargins.xsmall,
    paddingRight: globalMargins.xsmall,
  },
})

const containerStyles = styleSheetCreate({
  Custom: {},
  Danger: {...common, backgroundColor: globalColors.red},
  Primary: {...common, backgroundColor: globalColors.blue},
  PrimaryColoredBackgroundBlack: {...common, backgroundColor: globalColors.white},
  PrimaryColoredBackgroundBlue: {...common, backgroundColor: globalColors.white},
  PrimaryColoredBackgroundGreen: {...common, backgroundColor: globalColors.white},
  PrimaryColoredBackgroundPurple: {...common, backgroundColor: globalColors.white},
  PrimaryColoredBackgroundRed: {...common, backgroundColor: globalColors.white},
  PrimaryGreen: {...common, backgroundColor: globalColors.green},
  PrimaryGreenActive: platformStyles({
    common: {
      ...common,
      backgroundColor: globalColors.white,
      borderColor: globalColors.green,
      borderWidth: 2,
    },
    isElectron: {borderStyle: 'solid'},
  }),
  PrimaryPrivate: {...common, backgroundColor: globalColors.darkBlue2},
  Secondary: {...common, backgroundColor: globalColors.lightGrey2},
  SecondaryColoredBackgroundBlack: {...common, backgroundColor: globalColors.black_20},
  SecondaryColoredBackgroundBlue: {...common, backgroundColor: globalColors.black_20},
  SecondaryColoredBackgroundGreen: {...common, backgroundColor: globalColors.black_20},
  SecondaryColoredBackgroundPurple: {...common, backgroundColor: globalColors.black_20},
  SecondaryColoredBackgroundRed: {...common, backgroundColor: globalColors.black_20},
  SecondaryOnTerminal: {...common, backgroundColor: globalColors.blue_30},
  Wallet: {...common, backgroundColor: globalColors.purple2},
})

const labelStyles = styleSheetCreate({
  CustomLabel: {color: globalColors.black, textAlign: 'center'},
  DangerLabel: commonLabel,
  PrimaryColoredBackgroundLabelBlack: {...commonLabel, color: globalColors.black},
  PrimaryColoredBackgroundLabelBlue: {...commonLabel, color: globalColors.blue},
  PrimaryColoredBackgroundLabelGreen: {...commonLabel, color: globalColors.green},
  PrimaryColoredBackgroundLabelPurple: {...commonLabel, color: globalColors.purple},
  PrimaryColoredBackgroundLabelRed: {...commonLabel, color: globalColors.red},
  PrimaryGreenActiveLabel: {...commonLabel, color: globalColors.green},
  PrimaryGreenLabel: commonLabel,
  PrimaryLabel: commonLabel,
  PrimaryPrivateLabel: commonLabel,
  SecondaryColoredBackgroundLabel: {...commonLabel, color: globalColors.white},
  SecondaryColoredBackgroundLabelBlack: {...commonLabel, color: globalColors.white},
  SecondaryColoredBackgroundLabelBlue: {...commonLabel, color: globalColors.white},
  SecondaryColoredBackgroundLabelGreen: {...commonLabel, color: globalColors.white},
  SecondaryColoredBackgroundLabelPurple: {...commonLabel, color: globalColors.white},
  SecondaryColoredBackgroundLabelRed: {...commonLabel, color: globalColors.white},
  SecondaryLabel: {...commonLabel, color: globalColors.black},
  SecondaryLabelOnTerminal: {...commonLabel, color: globalColors.white},
  WalletLabel: commonLabel,
})

export default Button
