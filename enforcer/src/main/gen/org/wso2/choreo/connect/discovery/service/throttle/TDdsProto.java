// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/service/throtlle/tdds.proto

package org.wso2.choreo.connect.discovery.service.throttle;

public final class TDdsProto {
  private TDdsProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n*wso2/discovery/service/throtlle/tdds.p" +
      "roto\022\032discovery.service.throttle\032*envoy/" +
      "service/discovery/v3/discovery.proto2\213\002\n" +
      "\034ThrottleDataDiscoveryService\022w\n\022StreamT" +
      "hrottleData\022,.envoy.service.discovery.v3" +
      ".DiscoveryRequest\032-.envoy.service.discov" +
      "ery.v3.DiscoveryResponse\"\000(\0010\001\022r\n\021FetchT" +
      "hrottleData\022,.envoy.service.discovery.v3" +
      ".DiscoveryRequest\032-.envoy.service.discov" +
      "ery.v3.DiscoveryResponse\"\000B\214\001\n2org.wso2." +
      "choreo.connect.discovery.service.throttl" +
      "eB\tTDdsProtoP\001ZFgithub.com/envoyproxy/go" +
      "-control-plane/wso2/discovery/service/th" +
      "rottle\210\001\001b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          io.envoyproxy.envoy.service.discovery.v3.DiscoveryProto.getDescriptor(),
        });
    io.envoyproxy.envoy.service.discovery.v3.DiscoveryProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
