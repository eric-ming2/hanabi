// This file is @generated by prost-build.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Request {
    #[prost(string, tag = "1")]
    pub id: ::prost::alloc::string::String,
    #[prost(enumeration = "RequestType", tag = "2")]
    pub request_type: i32,
    #[prost(oneof = "request::Request", tags = "3, 4, 5, 6, 7, 8")]
    pub request: ::core::option::Option<request::Request>,
}
/// Nested message and enum types in `Request`.
pub mod request {
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum Request {
        #[prost(message, tag = "3")]
        InitConnection(super::InitConnectionRequest),
        #[prost(message, tag = "4")]
        Ready(super::ReadyRequest),
        #[prost(message, tag = "5")]
        StartGame(super::StartGameRequest),
        #[prost(message, tag = "6")]
        DiscardCard(super::DiscardCardRequest),
        #[prost(message, tag = "7")]
        PlayCard(super::PlayCardRequest),
        #[prost(message, tag = "8")]
        GiveHint(super::GiveHintRequest),
    }
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct InitConnectionRequest {
    #[prost(string, tag = "1")]
    pub id: ::prost::alloc::string::String,
    #[prost(string, tag = "2")]
    pub username: ::prost::alloc::string::String,
}
#[derive(Clone, Copy, PartialEq, ::prost::Message)]
pub struct ReadyRequest {}
#[derive(Clone, Copy, PartialEq, ::prost::Message)]
pub struct StartGameRequest {}
#[derive(Clone, Copy, PartialEq, ::prost::Message)]
pub struct DiscardCardRequest {
    #[prost(int32, tag = "1")]
    pub card_index: i32,
}
#[derive(Clone, Copy, PartialEq, ::prost::Message)]
pub struct PlayCardRequest {
    #[prost(int32, tag = "1")]
    pub card_index: i32,
}
#[derive(Clone, Copy, PartialEq, ::prost::Message)]
pub struct GiveHintRequest {
    #[prost(int32, tag = "1")]
    pub card_index: i32,
    #[prost(int32, tag = "2")]
    pub player_index: i32,
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum RequestType {
    InitConnection = 0,
    Ready = 1,
    StartGame = 2,
    DiscardCard = 3,
    PlayCard = 4,
    GiveHint = 5,
}
impl RequestType {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            Self::InitConnection => "INIT_CONNECTION",
            Self::Ready => "READY",
            Self::StartGame => "START_GAME",
            Self::DiscardCard => "DISCARD_CARD",
            Self::PlayCard => "PLAY_CARD",
            Self::GiveHint => "GIVE_HINT",
        }
    }
    /// Creates an enum from field names used in the ProtoBuf definition.
    pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
        match value {
            "INIT_CONNECTION" => Some(Self::InitConnection),
            "READY" => Some(Self::Ready),
            "START_GAME" => Some(Self::StartGame),
            "DISCARD_CARD" => Some(Self::DiscardCard),
            "PLAY_CARD" => Some(Self::PlayCard),
            "GIVE_HINT" => Some(Self::GiveHint),
            _ => None,
        }
    }
}
