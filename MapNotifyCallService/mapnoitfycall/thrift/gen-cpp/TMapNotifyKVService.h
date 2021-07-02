/**
 * Autogenerated by Thrift Compiler (0.11.0)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#ifndef TMapNotifyKVService_H
#define TMapNotifyKVService_H

#include <thrift/TDispatchProcessor.h>
#include <thrift/async/TConcurrentClientSyncInfo.h>
#include "i2skv_types.h"
#include "TDataService.h"

namespace OpenStars { namespace Common { namespace MapPhoneNumberPubkeyKV {

#ifdef _MSC_VER
  #pragma warning( push )
  #pragma warning (disable : 4250 ) //inheriting methods via dominance 
#endif

class TMapNotifyKVServiceIf : virtual public TDataServiceIf {
 public:
  virtual ~TMapNotifyKVServiceIf() {}
};

class TMapNotifyKVServiceIfFactory : virtual public TDataServiceIfFactory {
 public:
  typedef TMapNotifyKVServiceIf Handler;

  virtual ~TMapNotifyKVServiceIfFactory() {}

  virtual TMapNotifyKVServiceIf* getHandler(const ::apache::thrift::TConnectionInfo& connInfo) = 0;
  virtual void releaseHandler(TDataServiceIf* /* handler */) = 0;
};

class TMapNotifyKVServiceIfSingletonFactory : virtual public TMapNotifyKVServiceIfFactory {
 public:
  TMapNotifyKVServiceIfSingletonFactory(const ::apache::thrift::stdcxx::shared_ptr<TMapNotifyKVServiceIf>& iface) : iface_(iface) {}
  virtual ~TMapNotifyKVServiceIfSingletonFactory() {}

  virtual TMapNotifyKVServiceIf* getHandler(const ::apache::thrift::TConnectionInfo&) {
    return iface_.get();
  }
  virtual void releaseHandler(TDataServiceIf* /* handler */) {}

 protected:
  ::apache::thrift::stdcxx::shared_ptr<TMapNotifyKVServiceIf> iface_;
};

class TMapNotifyKVServiceNull : virtual public TMapNotifyKVServiceIf , virtual public TDataServiceNull {
 public:
  virtual ~TMapNotifyKVServiceNull() {}
};

class TMapNotifyKVServiceClient : virtual public TMapNotifyKVServiceIf, public TDataServiceClient {
 public:
  TMapNotifyKVServiceClient(apache::thrift::stdcxx::shared_ptr< ::apache::thrift::protocol::TProtocol> prot) :
    TDataServiceClient(prot, prot) {}
  TMapNotifyKVServiceClient(apache::thrift::stdcxx::shared_ptr< ::apache::thrift::protocol::TProtocol> iprot, apache::thrift::stdcxx::shared_ptr< ::apache::thrift::protocol::TProtocol> oprot) :    TDataServiceClient(iprot, oprot) {}
  apache::thrift::stdcxx::shared_ptr< ::apache::thrift::protocol::TProtocol> getInputProtocol() {
    return piprot_;
  }
  apache::thrift::stdcxx::shared_ptr< ::apache::thrift::protocol::TProtocol> getOutputProtocol() {
    return poprot_;
  }
};

class TMapNotifyKVServiceProcessor : public TDataServiceProcessor {
 protected:
  ::apache::thrift::stdcxx::shared_ptr<TMapNotifyKVServiceIf> iface_;
  virtual bool dispatchCall(::apache::thrift::protocol::TProtocol* iprot, ::apache::thrift::protocol::TProtocol* oprot, const std::string& fname, int32_t seqid, void* callContext);
 private:
  typedef  void (TMapNotifyKVServiceProcessor::*ProcessFunction)(int32_t, ::apache::thrift::protocol::TProtocol*, ::apache::thrift::protocol::TProtocol*, void*);
  typedef std::map<std::string, ProcessFunction> ProcessMap;
  ProcessMap processMap_;
 public:
  TMapNotifyKVServiceProcessor(::apache::thrift::stdcxx::shared_ptr<TMapNotifyKVServiceIf> iface) :
    TDataServiceProcessor(iface),
    iface_(iface) {
  }

  virtual ~TMapNotifyKVServiceProcessor() {}
};

class TMapNotifyKVServiceProcessorFactory : public ::apache::thrift::TProcessorFactory {
 public:
  TMapNotifyKVServiceProcessorFactory(const ::apache::thrift::stdcxx::shared_ptr< TMapNotifyKVServiceIfFactory >& handlerFactory) :
      handlerFactory_(handlerFactory) {}

  ::apache::thrift::stdcxx::shared_ptr< ::apache::thrift::TProcessor > getProcessor(const ::apache::thrift::TConnectionInfo& connInfo);

 protected:
  ::apache::thrift::stdcxx::shared_ptr< TMapNotifyKVServiceIfFactory > handlerFactory_;
};

class TMapNotifyKVServiceMultiface : virtual public TMapNotifyKVServiceIf, public TDataServiceMultiface {
 public:
  TMapNotifyKVServiceMultiface(std::vector<apache::thrift::stdcxx::shared_ptr<TMapNotifyKVServiceIf> >& ifaces) : ifaces_(ifaces) {
    std::vector<apache::thrift::stdcxx::shared_ptr<TMapNotifyKVServiceIf> >::iterator iter;
    for (iter = ifaces.begin(); iter != ifaces.end(); ++iter) {
      TDataServiceMultiface::add(*iter);
    }
  }
  virtual ~TMapNotifyKVServiceMultiface() {}
 protected:
  std::vector<apache::thrift::stdcxx::shared_ptr<TMapNotifyKVServiceIf> > ifaces_;
  TMapNotifyKVServiceMultiface() {}
  void add(::apache::thrift::stdcxx::shared_ptr<TMapNotifyKVServiceIf> iface) {
    TDataServiceMultiface::add(iface);
    ifaces_.push_back(iface);
  }
 public:
};

// The 'concurrent' client is a thread safe client that correctly handles
// out of order responses.  It is slower than the regular client, so should
// only be used when you need to share a connection among multiple threads
class TMapNotifyKVServiceConcurrentClient : virtual public TMapNotifyKVServiceIf, public TDataServiceConcurrentClient {
 public:
  TMapNotifyKVServiceConcurrentClient(apache::thrift::stdcxx::shared_ptr< ::apache::thrift::protocol::TProtocol> prot) :
    TDataServiceConcurrentClient(prot, prot) {}
  TMapNotifyKVServiceConcurrentClient(apache::thrift::stdcxx::shared_ptr< ::apache::thrift::protocol::TProtocol> iprot, apache::thrift::stdcxx::shared_ptr< ::apache::thrift::protocol::TProtocol> oprot) :    TDataServiceConcurrentClient(iprot, oprot) {}
  apache::thrift::stdcxx::shared_ptr< ::apache::thrift::protocol::TProtocol> getInputProtocol() {
    return piprot_;
  }
  apache::thrift::stdcxx::shared_ptr< ::apache::thrift::protocol::TProtocol> getOutputProtocol() {
    return poprot_;
  }
};

#ifdef _MSC_VER
  #pragma warning( pop )
#endif

}}} // namespace

#endif
