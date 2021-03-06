package main

import icon "github.com/wizenerd/icons"

// common mdl icon s
var names = []icon.Name{
	icon.ThreeDRotation,
	icon.AcUnit,
	icon.AccessAlarm,
	icon.AccessAlarms,
	icon.AccessTime,
	icon.Accessibility,
	icon.Accessible,
	icon.AccountBalance,
	icon.AccountBalanceWallet,
	icon.AccountBox,
	icon.AccountCircle,
	icon.Adb,
	icon.Add,
	icon.AddAPhoto,
	icon.AddAlarm,
	icon.AddAlert,
	icon.AddBox,
	icon.AddCircle,
	icon.AddCircleOutline,
	icon.AddLocation,
	icon.AddShoppingCart,
	icon.AddToPhotos,
	icon.AddToQueue,
	icon.Adjust,
	icon.AirlineSeatFlat,
	icon.AirlineSeatFlatAngled,
	icon.AirlineSeatIndividualSuite,
	icon.AirlineSeatLegroomExtra,
	icon.AirlineSeatLegroomNormal,
	icon.AirlineSeatLegroomReduced,
	icon.AirlineSeatReclineExtra,
	icon.AirlineSeatReclineNormal,
	icon.AirplanemodeActive,
	icon.AirplanemodeInactive,
	icon.Airplay,
	icon.AirportShuttle,
	icon.Alarm,
	icon.AlarmAdd,
	icon.AlarmOff,
	icon.AlarmOn,
	icon.Album,
	icon.AllInclusive,
	icon.AllOut,
	icon.Android,
	icon.Announcement,
	icon.Apps,
	icon.Archive,
	icon.ArrowBack,
	icon.ArrowDownward,
	icon.ArrowDropDown,
	icon.ArrowDropDownCircle,
	icon.ArrowDropUp,
	icon.ArrowForward,
	icon.ArrowUpward,
	icon.ArtTrack,
	icon.AspectRatio,
	icon.Assessment,
	icon.Assignment,
	icon.AssignmentInd,
	icon.AssignmentLate,
	icon.AssignmentReturn,
	icon.AssignmentReturned,
	icon.AssignmentTurnedIn,
	icon.Assistant,
	icon.AssistantPhoto,
	icon.AttachFile,
	icon.AttachMoney,
	icon.Attachment,
	icon.Audiotrack,
	icon.Autorenew,
	icon.AvTimer,
	icon.Backspace,
	icon.Backup,
	icon.BatteryAlert,
	icon.BatteryChargingFull,
	icon.BatteryFull,
	icon.BatteryStd,
	icon.BatteryUnknown,
	icon.BeachAccess,
	icon.Beenhere,
	icon.Block,
	icon.Bluetooth,
	icon.BluetoothAudio,
	icon.BluetoothConnected,
	icon.BluetoothDisabled,
	icon.BluetoothSearching,
	icon.BlurCircular,
	icon.BlurLinear,
	icon.BlurOff,
	icon.BlurOn,
	icon.Book,
	icon.Bookmark,
	icon.BookmarkBorder,
	icon.BorderAll,
	icon.BorderBottom,
	icon.BorderClear,
	icon.BorderColor,
	icon.BorderHorizontal,
	icon.BorderInner,
	icon.BorderLeft,
	icon.BorderOuter,
	icon.BorderRight,
	icon.BorderStyle,
	icon.BorderTop,
	icon.BorderVertical,
	icon.BrandingWatermark,
	icon.Brightness1,
	icon.Brightness2,
	icon.Brightness3,
	icon.Brightness4,
	icon.Brightness5,
	icon.Brightness6,
	icon.Brightness7,
	icon.BrightnessAuto,
	icon.BrightnessHigh,
	icon.BrightnessLow,
	icon.BrightnessMedium,
	icon.BrokenImage,
	icon.Brush,
	icon.BubbleChart,
	icon.BugReport,
	icon.Build,
	icon.BurstMode,
	icon.Business,
	icon.BusinessCenter,
	icon.Cached,
	icon.Cake,
	icon.Call,
	icon.CallEnd,
	icon.CallMade,
	icon.CallMerge,
	icon.CallMissed,
	icon.CallMissedOutgoing,
	icon.CallReceived,
	icon.CallSplit,
	icon.CallToAction,
	icon.Camera,
	icon.CameraAlt,
	icon.CameraEnhance,
	icon.CameraFront,
	icon.CameraRear,
	icon.CameraRoll,
	icon.Cancel,
	icon.CardGiftcard,
	icon.CardMembership,
	icon.CardTravel,
	icon.Casino,
	icon.Cast,
	icon.CastConnected,
	icon.CenterFocusStrong,
	icon.CenterFocusWeak,
	icon.ChangeHistory,
	icon.Chat,
	icon.ChatBubble,
	icon.ChatBubbleOutline,
	icon.Check,
	icon.CheckBox,
	icon.CheckBoxOutlineBlank,
	icon.CheckCircle,
	icon.ChevronLeft,
	icon.ChevronRight,
	icon.ChildCare,
	icon.ChildFriendly,
	icon.ChromeReaderMode,
	icon.Class,
	icon.Clear,
	icon.ClearAll,
	icon.Close,
	icon.ClosedCaption,
	icon.Cloud,
	icon.CloudCircle,
	icon.CloudDone,
	icon.CloudDownload,
	icon.CloudOff,
	icon.CloudQueue,
	icon.CloudUpload,
	icon.Code,
	icon.Collections,
	icon.CollectionsBookmark,
	icon.ColorLens,
	icon.Colorize,
	icon.Comment,
	icon.Compare,
	icon.CompareArrows,
	icon.Computer,
	icon.ConfirmationNumber,
	icon.ContactMail,
	icon.ContactPhone,
	icon.Contacts,
	icon.ContentCopy,
	icon.ContentCut,
	icon.ContentPaste,
	icon.ControlPoint,
	icon.ControlPointDuplicate,
	icon.Copyright,
	icon.Create,
	icon.CreateNewFolder,
	icon.CreditCard,
	icon.Crop,
	icon.Crop169,
	icon.Crop32,
	icon.Crop54,
	icon.Crop75,
	icon.CropDin,
	icon.CropFree,
	icon.CropLandscape,
	icon.CropOriginal,
	icon.CropPortrait,
	icon.CropRotate,
	icon.CropSquare,
	icon.Dashboard,
	icon.DataUsage,
	icon.DateRange,
	icon.Dehaze,
	icon.Delete,
	icon.DeleteForever,
	icon.DeleteSweep,
	icon.Description,
	icon.DesktopMac,
	icon.DesktopWindows,
	icon.Details,
	icon.DeveloperBoard,
	icon.DeveloperMode,
	icon.DeviceHub,
	icon.Devices,
	icon.DevicesOther,
	icon.DialerSip,
	icon.Dialpad,
	icon.Directions,
	icon.DirectionsBike,
	icon.DirectionsBoat,
	icon.DirectionsBus,
	icon.DirectionsCar,
	icon.DirectionsRailway,
	icon.DirectionsRun,
	icon.DirectionsSubway,
	icon.DirectionsTransit,
	icon.DirectionsWalk,
	icon.DiscFull,
	icon.Dns,
	icon.DoNotDisturb,
	icon.DoNotDisturbAlt,
	icon.DoNotDisturbOff,
	icon.DoNotDisturbOn,
	icon.Dock,
	icon.Domain,
	icon.Done,
	icon.DoneAll,
	icon.DonutLarge,
	icon.DonutSmall,
	icon.Drafts,
	icon.DragHandle,
	icon.DriveEta,
	icon.Dvr,
	icon.Edit,
	icon.EditLocation,
	icon.Eject,
	icon.Email,
	icon.EnhancedEncryption,
	icon.Equalizer,
	icon.Error,
	icon.ErrorOutline,
	icon.EuroSymbol,
	icon.EvStation,
	icon.Event,
	icon.EventAvailable,
	icon.EventBusy,
	icon.EventNote,
	icon.EventSeat,
	icon.ExitToApp,
	icon.ExpandLess,
	icon.ExpandMore,
	icon.Explicit,
	icon.Explore,
	icon.Exposure,
	icon.ExposureNeg1,
	icon.ExposureNeg2,
	icon.ExposurePlus1,
	icon.ExposurePlus2,
	icon.ExposureZero,
	icon.Extension,
	icon.Face,
	icon.FastForward,
	icon.FastRewind,
	icon.Favorite,
	icon.FavoriteBorder,
	icon.FeaturedPlayList,
	icon.FeaturedVideo,
	icon.Feedback,
	icon.FiberDvr,
	icon.FiberManualRecord,
	icon.FiberNew,
	icon.FiberPin,
	icon.FiberSmartRecord,
	icon.FileDownload,
	icon.FileUpload,
	icon.Filter,
	icon.Filter1,
	icon.Filter2,
	icon.Filter3,
	icon.Filter4,
	icon.Filter5,
	icon.Filter6,
	icon.Filter7,
	icon.Filter8,
	icon.Filter9,
	icon.Filter9Plus,
	icon.FilterBAndW,
	icon.FilterCenterFocus,
	icon.FilterDrama,
	icon.FilterFrames,
	icon.FilterHdr,
	icon.FilterList,
	icon.FilterNone,
	icon.FilterTiltShift,
	icon.FilterVintage,
	icon.FindInPage,
	icon.FindReplace,
	icon.Fingerprint,
	icon.FirstPage,
	icon.FitnessCenter,
	icon.Flag,
	icon.Flare,
	icon.FlashAuto,
	icon.FlashOff,
	icon.FlashOn,
	icon.Flight,
	icon.FlightLand,
	icon.FlightTakeoff,
	icon.Flip,
	icon.FlipToBack,
	icon.FlipToFront,
	icon.Folder,
	icon.FolderOpen,
	icon.FolderShared,
	icon.FolderSpecial,
	icon.FontDownload,
	icon.FormatAlignCenter,
	icon.FormatAlignJustify,
	icon.FormatAlignLeft,
	icon.FormatAlignRight,
	icon.FormatBold,
	icon.FormatClear,
	icon.FormatColorFill,
	icon.FormatColorReset,
	icon.FormatColorText,
	icon.FormatIndentDecrease,
	icon.FormatIndentIncrease,
	icon.FormatItalic,
	icon.FormatLineSpacing,
	icon.FormatListBulleted,
	icon.FormatListNumbered,
	icon.FormatPaint,
	icon.FormatQuote,
	icon.FormatShapes,
	icon.FormatSize,
	icon.FormatStrikethrough,
	icon.FormatTextdirectionLToR,
	icon.FormatTextdirectionRToL,
	icon.FormatUnderlined,
	icon.Forum,
	icon.Forward,
	icon.Forward10,
	icon.Forward30,
	icon.Forward5,
	icon.FreeBreakfast,
	icon.Fullscreen,
	icon.FullscreenExit,
	icon.Functions,
	icon.GTranslate,
	icon.Gamepad,
	icon.Games,
	icon.Gavel,
	icon.Gesture,
	icon.GetApp,
	icon.Gif,
	icon.GolfCourse,
	icon.GpsFixed,
	icon.GpsNotFixed,
	icon.GpsOff,
	icon.Grade,
	icon.Gradient,
	icon.Grain,
	icon.GraphicEq,
	icon.GridOff,
	icon.GridOn,
	icon.Group,
	icon.GroupAdd,
	icon.GroupWork,
	icon.Hd,
	icon.HdrOff,
	icon.HdrOn,
	icon.HdrStrong,
	icon.HdrWeak,
	icon.Headset,
	icon.HeadsetMic,
	icon.Healing,
	icon.Hearing,
	icon.Help,
	icon.HelpOutline,
	icon.HighQuality,
	icon.Highlight,
	icon.HighlightOff,
	icon.History,
	icon.Home,
	icon.HotTub,
	icon.Hotel,
	icon.HourglassEmpty,
	icon.HourglassFull,
	icon.Http,
	icon.Https,
	icon.Image,
	icon.ImageAspectRatio,
	icon.ImportContacts,
	icon.ImportExport,
	icon.ImportantDevices,
	icon.Inbox,
	icon.IndeterminateCheckBox,
	icon.Info,
	icon.InfoOutline,
	icon.Input,
	icon.InsertChart,
	icon.InsertComment,
	icon.InsertDriveFile,
	icon.InsertEmoticon,
	icon.InsertInvitation,
	icon.InsertLink,
	icon.InsertPhoto,
	icon.InvertColors,
	icon.InvertColorsOff,
	icon.Iso,
	icon.Keyboard,
	icon.KeyboardArrowDown,
	icon.KeyboardArrowLeft,
	icon.KeyboardArrowRight,
	icon.KeyboardArrowUp,
	icon.KeyboardBackspace,
	icon.KeyboardCapslock,
	icon.KeyboardHide,
	icon.KeyboardReturn,
	icon.KeyboardTab,
	icon.KeyboardVoice,
	icon.Kitchen,
	icon.Label,
	icon.LabelOutline,
	icon.Landscape,
	icon.Language,
	icon.Laptop,
	icon.LaptopChromebook,
	icon.LaptopMac,
	icon.LaptopWindows,
	icon.LastPage,
	icon.Launch,
	icon.Layers,
	icon.LayersClear,
	icon.LeakAdd,
	icon.LeakRemove,
	icon.Lens,
	icon.LibraryAdd,
	icon.LibraryBooks,
	icon.LibraryMusic,
	icon.LightbulbOutline,
	icon.LineStyle,
	icon.LineWeight,
	icon.LinearScale,
	icon.Link,
	icon.LinkedCamera,
	icon.List,
	icon.LiveHelp,
	icon.LiveTv,
	icon.LocalActivity,
	icon.LocalAirport,
	icon.LocalAtm,
	icon.LocalBar,
	icon.LocalCafe,
	icon.LocalCarWash,
	icon.LocalConvenienceStore,
	icon.LocalDining,
	icon.LocalDrink,
	icon.LocalFlorist,
	icon.LocalGasStation,
	icon.LocalGroceryStore,
	icon.LocalHospital,
	icon.LocalHotel,
	icon.LocalLaundryService,
	icon.LocalLibrary,
	icon.LocalMall,
	icon.LocalMovies,
	icon.LocalOffer,
	icon.LocalParking,
	icon.LocalPharmacy,
	icon.LocalPhone,
	icon.LocalPizza,
	icon.LocalPlay,
	icon.LocalPostOffice,
	icon.LocalPrintshop,
	icon.LocalSee,
	icon.LocalShipping,
	icon.LocalTaxi,
	icon.LocationCity,
	icon.LocationDisabled,
	icon.LocationOff,
	icon.LocationOn,
	icon.LocationSearching,
	icon.Lock,
	icon.LockOpen,
	icon.LockOutline,
	icon.Looks,
	icon.Looks3,
	icon.Looks4,
	icon.Looks5,
	icon.Looks6,
	icon.LooksOne,
	icon.LooksTwo,
	icon.Loop,
	icon.Loupe,
	icon.LowPriority,
	icon.Loyalty,
	icon.Mail,
	icon.MailOutline,
	icon.Map,
	icon.Markunread,
	icon.MarkunreadMailbox,
	icon.Memory,
	icon.Menu,
	icon.MergeType,
	icon.Message,
	icon.Mic,
	icon.MicNone,
	icon.MicOff,
	icon.Mms,
	icon.ModeComment,
	icon.ModeEdit,
	icon.MonetizationOn,
	icon.MoneyOff,
	icon.MonochromePhotos,
	icon.Mood,
	icon.MoodBad,
	icon.More,
	icon.MoreHoriz,
	icon.MoreVert,
	icon.Motorcycle,
	icon.Mouse,
	icon.MoveToInbox,
	icon.Movie,
	icon.MovieCreation,
	icon.MovieFilter,
	icon.MultilineChart,
	icon.MusicNote,
	icon.MusicVideo,
	icon.MyLocation,
	icon.Nature,
	icon.NaturePeople,
	icon.NavigateBefore,
	icon.NavigateNext,
	icon.Navigation,
	icon.NearMe,
	icon.NetworkCell,
	icon.NetworkCheck,
	icon.NetworkLocked,
	icon.NetworkWifi,
	icon.NewReleases,
	icon.NextWeek,
	icon.Nfc,
	icon.NoEncryption,
	icon.NoSim,
	icon.NotInterested,
	icon.Note,
	icon.NoteAdd,
	icon.Notifications,
	icon.NotificationsActive,
	icon.NotificationsNone,
	icon.NotificationsOff,
	icon.NotificationsPaused,
	icon.OfflinePin,
	icon.OndemandVideo,
	icon.Opacity,
	icon.OpenInBrowser,
	icon.OpenInNew,
	icon.OpenWith,
	icon.Pages,
	icon.Pageview,
	icon.Palette,
	icon.PanTool,
	icon.Panorama,
	icon.PanoramaFishEye,
	icon.PanoramaHorizontal,
	icon.PanoramaVertical,
	icon.PanoramaWideAngle,
	icon.PartyMode,
	icon.Pause,
	icon.PauseCircleFilled,
	icon.PauseCircleOutline,
	icon.Payment,
	icon.People,
	icon.PeopleOutline,
	icon.PermCameraMic,
	icon.PermContactCalendar,
	icon.PermDataSetting,
	icon.PermDeviceInformation,
	icon.PermIdentity,
	icon.PermMedia,
	icon.PermPhoneMsg,
	icon.PermScanWifi,
	icon.Person,
	icon.PersonAdd,
	icon.PersonOutline,
	icon.PersonPin,
	icon.PersonPinCircle,
	icon.PersonalVideo,
	icon.Pets,
	icon.Phone,
	icon.PhoneAndroid,
	icon.PhoneBluetoothSpeaker,
	icon.PhoneForwarded,
	icon.PhoneInTalk,
	icon.PhoneIphone,
	icon.PhoneLocked,
	icon.PhoneMissed,
	icon.PhonePaused,
	icon.Phonelink,
	icon.PhonelinkErase,
	icon.PhonelinkLock,
	icon.PhonelinkOff,
	icon.PhonelinkRing,
	icon.PhonelinkSetup,
	icon.Photo,
	icon.PhotoAlbum,
	icon.PhotoCamera,
	icon.PhotoFilter,
	icon.PhotoLibrary,
	icon.PhotoSizeSelectActual,
	icon.PhotoSizeSelectLarge,
	icon.PhotoSizeSelectSmall,
	icon.PictureAsPdf,
	icon.PictureInPicture,
	icon.PictureInPictureAlt,
	icon.PieChart,
	icon.PieChartOutlined,
	icon.PinDrop,
	icon.Place,
	icon.PlayArrow,
	icon.PlayCircleFilled,
	icon.PlayCircleOutline,
	icon.PlayForWork,
	icon.PlaylistAdd,
	icon.PlaylistAddCheck,
	icon.PlaylistPlay,
	icon.PlusOne,
	icon.Poll,
	icon.Polymer,
	icon.Pool,
	icon.PortableWifiOff,
	icon.Portrait,
	icon.Power,
	icon.PowerInput,
	icon.PowerSettingsNew,
	icon.PregnantWoman,
	icon.PresentToAll,
	icon.Print,
	icon.PriorityHigh,
	icon.Public,
	icon.Publish,
	icon.QueryBuilder,
	icon.QuestionAnswer,
	icon.Queue,
	icon.QueueMusic,
	icon.QueuePlayNext,
	icon.Radio,
	icon.RadioButtonChecked,
	icon.RadioButtonUnchecked,
	icon.RateReview,
	icon.Receipt,
	icon.RecentActors,
	icon.RecordVoiceOver,
	icon.Redeem,
	icon.Redo,
	icon.Refresh,
	icon.Remove,
	icon.RemoveCircle,
	icon.RemoveCircleOutline,
	icon.RemoveFromQueue,
	icon.RemoveRedEye,
	icon.RemoveShoppingCart,
	icon.Reorder,
	icon.Repeat,
	icon.RepeatOne,
	icon.Replay,
	icon.Replay10,
	icon.Replay30,
	icon.Replay5,
	icon.Reply,
	icon.ReplyAll,
	icon.Report,
	icon.ReportProblem,
	icon.Restaurant,
	icon.RestaurantMenu,
	icon.Restore,
	icon.RestorePage,
	icon.RingVolume,
	icon.Room,
	icon.RoomService,
	icon.Rotate90DegreesCcw,
	icon.RotateLeft,
	icon.RotateRight,
	icon.RoundedCorner,
	icon.Router,
	icon.Rowing,
	icon.RssFeed,
	icon.RvHookup,
	icon.Satellite,
	icon.Save,
	icon.Scanner,
	icon.Schedule,
	icon.School,
	icon.ScreenLockLandscape,
	icon.ScreenLockPortrait,
	icon.ScreenLockRotation,
	icon.ScreenRotation,
	icon.ScreenShare,
	icon.SdCard,
	icon.SdStorage,
	icon.Search,
	icon.Security,
	icon.SelectAll,
	icon.Send,
	icon.SentimentDissatisfied,
	icon.SentimentNeutral,
	icon.SentimentSatisfied,
	icon.SentimentVeryDissatisfied,
	icon.SentimentVerySatisfied,
	icon.Settings,
	icon.SettingsApplications,
	icon.SettingsBackupRestore,
	icon.SettingsBluetooth,
	icon.SettingsBrightness,
	icon.SettingsCell,
	icon.SettingsEthernet,
	icon.SettingsInputAntenna,
	icon.SettingsInputComponent,
	icon.SettingsInputComposite,
	icon.SettingsInputHdmi,
	icon.SettingsInputSvideo,
	icon.SettingsOverscan,
	icon.SettingsPhone,
	icon.SettingsPower,
	icon.SettingsRemote,
	icon.SettingsSystemDaydream,
	icon.SettingsVoice,
	icon.Share,
	icon.Shop,
	icon.ShopTwo,
	icon.ShoppingBasket,
	icon.ShoppingCart,
	icon.ShortText,
	icon.ShowChart,
	icon.Shuffle,
	icon.SignalCellular4Bar,
	icon.SignalCellularConnectedNoInternet4Bar,
	icon.SignalCellularNoSim,
	icon.SignalCellularNull,
	icon.SignalCellularOff,
	icon.SignalWifi4Bar,
	icon.SignalWifi4BarLock,
	icon.SignalWifiOff,
	icon.SimCard,
	icon.SimCardAlert,
	icon.SkipNext,
	icon.SkipPrevious,
	icon.Slideshow,
	icon.SlowMotionVideo,
	icon.Smartphone,
	icon.SmokeFree,
	icon.SmokingRooms,
	icon.Sms,
	icon.SmsFailed,
	icon.Snooze,
	icon.Sort,
	icon.SortByAlpha,
	icon.Spa,
	icon.SpaceBar,
	icon.Speaker,
	icon.SpeakerGroup,
	icon.SpeakerNotes,
	icon.SpeakerNotesOff,
	icon.SpeakerPhone,
	icon.Spellcheck,
	icon.Star,
	icon.StarBorder,
	icon.StarHalf,
	icon.Stars,
	icon.StayCurrentLandscape,
	icon.StayCurrentPortrait,
	icon.StayPrimaryLandscape,
	icon.StayPrimaryPortrait,
	icon.Stop,
	icon.StopScreenShare,
	icon.Storage,
	icon.Store,
	icon.StoreMallDirectory,
	icon.Straighten,
	icon.Streetview,
	icon.StrikethroughS,
	icon.Style,
	icon.SubdirectoryArrowLeft,
	icon.SubdirectoryArrowRight,
	icon.Subject,
	icon.Subscriptions,
	icon.Subtitles,
	icon.Subway,
	icon.SupervisorAccount,
	icon.SurroundSound,
	icon.SwapCalls,
	icon.SwapHoriz,
	icon.SwapVert,
	icon.SwapVerticalCircle,
	icon.SwitchCamera,
	icon.SwitchVideo,
	icon.Sync,
	icon.SyncDisabled,
	icon.SyncProblem,
	icon.SystemUpdate,
	icon.SystemUpdateAlt,
	icon.Tab,
	icon.TabUnselected,
	icon.Tablet,
	icon.TabletAndroid,
	icon.TabletMac,
	icon.TagFaces,
	icon.TapAndPlay,
	icon.Terrain,
	icon.TextFields,
	icon.TextFormat,
	icon.Textsms,
	icon.Texture,
	icon.Theaters,
	icon.ThumbDown,
	icon.ThumbUp,
	icon.ThumbsUpDown,
	icon.TimeToLeave,
	icon.Timelapse,
	icon.Timeline,
	icon.Timer,
	icon.Timer10,
	icon.Timer3,
	icon.TimerOff,
	icon.Title,
	icon.Toc,
	icon.Today,
	icon.Toll,
	icon.Tonality,
	icon.TouchApp,
	icon.Toys,
	icon.TrackChanges,
	icon.Traffic,
	icon.Train,
	icon.Tram,
	icon.TransferWithinAStation,
	icon.Transform,
	icon.Translate,
	icon.TrendingDown,
	icon.TrendingFlat,
	icon.TrendingUp,
	icon.Tune,
	icon.TurnedIn,
	icon.TurnedInNot,
	icon.Tv,
	icon.Unarchive,
	icon.Undo,
	icon.UnfoldLess,
	icon.UnfoldMore,
	icon.Update,
	icon.Usb,
	icon.VerifiedUser,
	icon.VerticalAlignBottom,
	icon.VerticalAlignCenter,
	icon.VerticalAlignTop,
	icon.Vibration,
	icon.VideoCall,
	icon.VideoLabel,
	icon.VideoLibrary,
	icon.Videocam,
	icon.VideocamOff,
	icon.VideogameAsset,
	icon.ViewAgenda,
	icon.ViewArray,
	icon.ViewCarousel,
	icon.ViewColumn,
	icon.ViewComfy,
	icon.ViewCompact,
	icon.ViewDay,
	icon.ViewHeadline,
	icon.ViewList,
	icon.ViewModule,
	icon.ViewQuilt,
	icon.ViewStream,
	icon.ViewWeek,
	icon.Vignette,
	icon.Visibility,
	icon.VisibilityOff,
	icon.VoiceChat,
	icon.Voicemail,
	icon.VolumeDown,
	icon.VolumeMute,
	icon.VolumeOff,
	icon.VolumeUp,
	icon.VpnKey,
	icon.VpnLock,
	icon.Wallpaper,
	icon.Warning,
	icon.Watch,
	icon.WatchLater,
	icon.WbAuto,
	icon.WbCloudy,
	icon.WbIncandescent,
	icon.WbIridescent,
	icon.WbSunny,
	icon.Wc,
	icon.Web,
	icon.WebAsset,
	icon.Weekend,
	icon.Whatshot,
	icon.Widgets,
	icon.Wifi,
	icon.WifiLock,
	icon.WifiTethering,
	icon.Work,
	icon.WrapText,
	icon.YoutubeSearchedFor,
	icon.ZoomIn,
	icon.ZoomOut,
	icon.ZoomOutMap,
}
